# Sync.Map

在 Go1.9 之前，Go 自带的 Map 不是并发安全的，因此我们需要自己再封装一层，给 Map 加上把读写锁，例如：

```go
type MapWithLock struct {
    sync.RWMutex
    M map[string]Kline
}
```

用 MapWithLock 的读写锁去控制 map 的并发安全，在 Go1.9 发布后，它有了一个新的特性，那就是 sync.map，它是原生支持并发安全的 map。

# 数据操作

# 内部优化

sync.Map 的实现和优化的点在于空间换时间。 通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。使用只读数据(read)，避免读写冲突。动态调整，miss 次数多了之后，将 dirty 数据提升为 read。double-checking。 延迟删除。 删除一个键值只是打标记，只有在提升 dirty 的时候才清理删除的数据。 优先从 read 读取、更新、删除，因为对 read 的读取不需要锁。

```go
type Map struct {
    // 当涉及到dirty数据的操作的时候，需要使用这个锁
    mu Mutex

    // 一个只读的数据结构，因为只读，所以不会有读写冲突。所以从这个数据中读取总是安全的。
    // 实际上，实际也会更新这个数据的entries,如果entry是未删除的(unexpunged), 并不需要加锁。如果entry已经被删除了，需要加锁，以便更新dirty数据。
    read atomic.Value // readOnly

    // dirty数据包含当前的map包含的entries,它包含最新的entries(包括read中未删除的数据,虽有冗余，但是提升dirty字段为read的时候非常快，不用一个一个的复制，而是直接将这个数据结构作为read字段的一部分),有些数据还可能没有移动到read字段中。
    // 对于dirty的操作需要加锁，因为对它的操作可能会有读写竞争。
    // 当dirty为空的时候， 比如初始化或者刚提升完，下一次的写操作会复制read字段中未删除的数据到这个数据中。
    dirty map[interface{}]*entry

    // 当从Map中读取entry的时候，如果read中不包含这个entry,会尝试从dirty中读取，这个时候会将misses加一，
    // 当misses累积到 dirty的长度的时候， 就会将dirty提升为read,避免从dirty中miss太多次。因为操作dirty需要加锁。
    misses int
}
```

它的数据结构很简单，值包含四个字段：read、mu、dirty、misses。readOnly.m 和 Map.dirty 存储的值类型是 `*entry`,它包含一个指针 p, 指向用户存储的 value 值。

```go
type entry struct {
    p unsafe.Pointer // *interface{}
}
```

p 通常有三种类型的值:

- nil: entry 已被删除了，并且 m.dirty 为 nil

- expunged: entry 已被删除了，并且 m.dirty 不为 nil，而且这个 entry 不存在于 m.dirty 中

- 其它： entry 是一个正常的值

它使用了冗余的数据结构 read、dirty。dirty 中会包含 read 中为删除的 entries，新增加的 entries 会加入到 dirty 中。 read 的数据结构是：

```go
type readOnly struct {
    m       map[interface{}]*entry
    amended bool // 如果Map.dirty有些数据不在中的时候，这个值为true
}
```

amended 指明 Map.dirty 中有 readOnly.m 未包含的数据，所以如果从 Map.read 找不到数据的话，还要进一步到 Map.dirty 中查找。对 Map.read 的修改是通过原子操作进行的，虽然 read 和 dirty 有冗余数据，但这些数据是通过指针指向同一个数据，所以尽管 Map 的 value 会很大，但是冗余的空间占用还是有限的。

## Load 方法

Load 方法 Load 方法，提供一个键 key,查找对应的值 value,如果不存在，通过 ok 反映：

```go
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
    // 1.首先从m.read中得到只读readOnly,从它的map中查找，不需要加锁
    read, _ := m.read.Load().(readOnly)
    e, ok := read.m[key]
    // 2. 如果没找到，并且m.dirty中有新数据，需要从m.dirty查找，这个时候需要加锁
    if !ok && read.amended {
        m.mu.Lock()
        // 双检查，避免加锁的时候m.dirty提升为m.read,这个时候m.read可能被替换了。
        read, _ = m.read.Load().(readOnly)
        e, ok = read.m[key]
        // 如果m.read中还是不存在，并且m.dirty中有新数据
        if !ok && read.amended {
            // 从m.dirty查找
            e, ok = m.dirty[key]
            // 不管m.dirty中存不存在，都将misses计数加一
            // missLocked()中满足条件后就会提升m.dirty
            m.missLocked()
        }
        m.mu.Unlock()
    }
    if !ok {
        return nil, false
    }
    return e.load()
}
```

这里会先从 m.read 中加载，不存在的情况下，并且 m.dirty 中有新数据，加锁，然后从 m.dirty 中加载。其次是这里使用了双检查的处理，因为在下面的两个语句中，这两行语句并不是一个原子操作。

```go
if !ok && read.amended {
        m.mu.Lock()
```

当第一句执行的时候条件满足，但是在加锁之前，m.dirty 可能被提升为 m.read,所以加锁后还得再检查 m.read，后续的方法中都使用了这个方法。如果我们查询的键值正好存在于 m.read 中，无须加锁，直接返回，理论上性能优异。即使不存在于 m.read 中，经过 miss 几次之后，m.dirty 会被提升为 m.read，又会从 m.read 中查找。所以对于更新／增加较少，加载存在的 key 很多的 case,性能基本和无锁的 map 类似。 接着我们看下如何 m.dirty 是如何被提升的。 missLocked 方法中可能会将 m.dirty 提升。

```go
func (m *Map) missLocked() {
    m.misses++
    if m.misses < len(m.dirty) {
        return
    }
    m.read.Store(readOnly{m: m.dirty})
    m.dirty = nil
    m.misses = 0
}
```

上面的最后三行代码就是提升 m.dirty 的，很简单的将 m.dirty 作为 readOnly 的 m 字段，原子更新 m.read。提升后 m.dirty、m.misses 重置， 并且 m.read.amended 为 false。

## Store

Store 方法是更新或者新增一个 entry。

```go
func (m *Map) Store(key, value interface{}) {
    // 如果m.read存在这个键，并且这个entry没有被标记删除，尝试直接存储。
    // 因为m.dirty也指向这个entry,所以m.dirty也保持最新的entry。
    read, _ := m.read.Load().(readOnly)
    if e, ok := read.m[key]; ok && e.tryStore(&value) {
        return
    }
    // 如果`m.read`不存在或者已经被标记删除
    m.mu.Lock()
    read, _ = m.read.Load().(readOnly)
    if e, ok := read.m[key]; ok {
        if e.unexpungeLocked() { //标记成未被删除
            m.dirty[key] = e //m.dirty中不存在这个键，所以加入m.dirty
        }
        e.storeLocked(&value) //更新
    } else if e, ok := m.dirty[key]; ok { // m.dirty存在这个键，更新
        e.storeLocked(&value)
    } else { //新键值
        if !read.amended { //m.dirty中没有新的数据，往m.dirty中增加第一个新键
            m.dirtyLocked() //从m.read中复制未删除的数据
            m.read.Store(readOnly{m: read.m, amended: true})
        }
        m.dirty[key] = newEntry(value) //将这个entry加入到m.dirty中
    }
    m.mu.Unlock()
}
func (m *Map) dirtyLocked() {
    if m.dirty != nil {
        return
    }
    read, _ := m.read.Load().(readOnly)
    m.dirty = make(map[interface{}]*entry, len(read.m))
    for k, e := range read.m {
        if !e.tryExpungeLocked() {
            m.dirty[k] = e
        }
    }
}
func (e *entry) tryExpungeLocked() (isExpunged bool) {
    p := atomic.LoadPointer(&e.p)
    for p == nil {
        // 将已经删除标记为nil的数据标记为expunged
        if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
            return true
        }
        p = atomic.LoadPointer(&e.p)
    }
    return p == expunged
}
```

通常是先从操作 m.read 开始的，如果不满足条件再加锁，然后操作 m.dirty。Store 方法可能会在某种情况下(初始化或者 m.dirty 刚被提升后)从 m.read 中复制数据，如果这个时候 m.read 中数据量非常大，可能会影响性能。

## Delete

Delete 方法用来删除一个键值。

```go
func (m *Map) Delete(key interface{}) {
    read, _ := m.read.Load().(readOnly)
    e, ok := read.m[key]
    if !ok && read.amended {
        m.mu.Lock()
        read, _ = m.read.Load().(readOnly)
        e, ok = read.m[key]
        if !ok && read.amended {
            delete(m.dirty, key)
        }
        m.mu.Unlock()
    }
    if ok {
        e.delete()
    }
}
```

这里的删除操作还是从 m.read 中开始， 如果这个 entry 不存在于 m.read 中，并且 m.dirty 中有新数据，则加锁尝试从 m.dirty 中删除。此外需要双检查的，从 m.dirty 中直接删除即可，就当它没存在过，但是如果是从 m.read 中删除，并不会直接删除，而是打标记：

```go
func (e *entry) delete() (hadValue bool) {
    for {
        p := atomic.LoadPointer(&e.p)
        // 已标记为删除
        if p == nil || p == expunged {
            return false
        }
        // 原子操作，e.p标记为nil
        if atomic.CompareAndSwapPointer(&e.p, p, nil) {
            return true
        }
    }
}
```

## Range

因为 for ... range map 是内建的语言特性，所以没有办法使用 for range 遍历 sync.Map, 但是可以使用它的 Range 方法，通过回调的方式遍历。

```go
func (m *Map) Range(f func(key, value interface{}) bool) {
    read, _ := m.read.Load().(readOnly)
    // 如果m.dirty中有新数据，则提升m.dirty,然后在遍历
    if read.amended {
        //提升m.dirty
        m.mu.Lock()
        read, _ = m.read.Load().(readOnly) //双检查
        if read.amended {
            read = readOnly{m: m.dirty}
            m.read.Store(read)
            m.dirty = nil
            m.misses = 0
        }
        m.mu.Unlock()
    }
    // 遍历, for range是安全的
    for k, e := range read.m {
        v, ok := e.load()
        if !ok {
            continue
        }
        if !f(k, v) {
            break
        }
    }
}
```

Range 方法调用前可能会做一个 m.dirty 的提升，不过提升 m.dirty 不是一个耗时的操作。
