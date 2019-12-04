# Sync.Map

在 Go1.9 之前，Go 自带的 Map 不是并发安全的，因此我们需要自己再封装一层，给 Map 加上把读写锁，例如：

```go
type MapWithLock struct {
    sync.RWMutex
    M map[string]Kline
}
```

用 MapWithLock 的读写锁去控制 map 的并发安全，在 Go1.9 发布后，它有了一个新的特性，那就是 sync.map，它是原生支持并发安全的 map。

# 内部优化

sync.Map 的实现和优化的点在于空间换时间。 通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。使用只读数据(read)，避免读写冲突。动态调整，miss 次数多了之后，将 dirty 数据提升为 read。double-checking。 延迟删除。 删除一个键值只是打标记，只有在提升 dirty 的时候才清理删除的数据。 优先从 read 读取、更新、删除，因为对 read 的读取不需要锁。

```go
type Map struct {
    // 当涉及到dirty数据的操作的时候，需要使用这个锁
    mu Mutex

    // 一个只读的数据结构，因为只读，所以不会有读写冲突。
    // 所以从这个数据中读取总是安全的。
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
