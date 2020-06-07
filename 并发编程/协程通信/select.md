# select

Golang 的 select 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。在执行 select 语句的时候，运行时系统会自上而下地判断每个 case 中的发送或接收操作是否可以被立即执行(当前 Goroutine 不会因此操作而被阻塞)。select 的用法与 switch 非常类似，由 select 开始一个新的选择块，每个选择条件由 case 语句来描述。

```go
var c1, c2 <-chan interface{}
var c3 chan<- interface{}
select {
	case <-c1:
		// Do something
	case <-c2:
		// Do something
	case c3 <- struct{}{}:
		// Do something
}
```

跟 switch 相同的是，select 代码块也包含一系列 case 分支。跟 switch 不同的是，case 分支不会被顺序测试，如果没有任何分支的条件可供满足，select 会一直等待直到某个 case 语句完成。所有通道的读取和写入都被同时考虑，以查看它们中的任何一个是否准备好：如果没有任何通道准备就绪，则整个 select 语句将会阻塞。当一个通道准备好时，该操作将继续，并执行相应的语句。我们来看一个简单的例子：

```go
start := time.Now()
c := make(chan interface{})
go func() {
	time.Sleep(5 * time.Second)
	close(c) // 5 秒后关闭通道
}()

fmt.Println("Blocking on read...")
select {
case <-c: // 尝试读取通道。注意，尽管我们可以不使用select语句而直接使用<-c，但我们的目的是为了展示select语句。
	fmt.Printf("Unblocked %v later.\n", time.Since(start))
}

// Blocking on read...
// Unblocked 5s later.
```

通道是将 Goroutine 的粘合剂，select 语句是通道的粘合剂。后者让我们能够在项目中组合通道以形成更大的抽象来解决实际中遇到的问题。我们可以在单个函数或类型定义中找到将本地通道绑定在一起的 select 语句，也可以在全局范围找到连接系统级别两个或多个组件的使用范例。除了连接组件外，在程序中的关键部分，select 语句还可以帮助你安全地将通道与业务层面的概念（如取消，超时，等待和默认值）结合在一起。

# select 顺序

对于空的 select 语句，程序会被阻塞，准确的说是当前协程被阻塞，同时 Golang 自带死锁检测机制，当发现当前协程再也没有机会被唤醒时，则会 panic。所以上述程序会 panic。

```go
package main

func main() {
    select {
    }
}
```

select 中各个 case 执行顺序是随机的，如果某个 case 中的 channel 已经 ready，则执行相应的语句并退出 select 流程，如果所有 case 中的 channel 都未 ready，则执行 default 中的语句然后退出 select 流程。另外，由于启动的协程和 select 语句并不能保证执行顺序，所以也有可能 select 执行时协程还未向 channel 中写入数据，所以 select 直接执行 default 语句并退出。所以，以下三种输出都有可能：

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    chan1 := make(chan int)
    chan2 := make(chan int)

    go func() {
        chan1 <- 1
        time.Sleep(5 * time.Second)
    }()

    go func() {
        chan2 <- 1
        time.Sleep(5 * time.Second)
    }()

    select {
    case <-chan1:
        fmt.Println("chan1 ready.")
    case <-chan2:
        fmt.Println("chan2 ready.")
    default:
        fmt.Println("default")
    }

    fmt.Println("main exit.")
}
```

select 会按照随机的顺序检测各 case 语句中 channel 是否 ready，如果某个 case 中的 channel 已经 ready 则执行相应的 case 语句然后退出 select 流程，如果所有的 channel 都未 ready 且没有 default 的话，则会阻塞等待各个 channel。所以上述程序会一直阻塞。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    chan1 := make(chan int)
    chan2 := make(chan int)

    writeFlag := false
    go func() {
        for {
            if writeFlag {
                chan1 <- 1
            }
            time.Sleep(time.Second)
        }
    }()

    go func() {
        for {
            if writeFlag {
                chan2 <- 1
            }
            time.Sleep(time.Second)
        }
    }()

    select {
    case <-chan1:
        fmt.Println("chan1 ready.")
    case <-chan2:
        fmt.Println("chan2 ready.")
    }

    fmt.Println("main exit.")
}
```

程序中声明两个 channel，分别为 chan1 和 chan2，依次启动两个协程，协程会判断一个 bool 类型的变量 writeFlag 来决定是否要向 channel 中写入数据，由于 writeFlag 永远为 false，所以实际上协程什么也没做。select 语句两个 case 分别检测 chan1 和 chan2 是否可读，这个 select 语句不包含 default 语句。

总结而言：

- select 语句中除 default 外，每个 case 操作一个 channel，要么读要么写
- select 语句中除 default 外，各 case 执行顺序是随机的
- select 语句中如果没有 default 语句，则会阻塞等待任一 case
- select 语句中读操作要判断是否成功读取，关闭的 channel 也可以读取

# select 的典型应用

## 定时器

```go
func main() {

   tickTimer := time.NewTicker(1 * time.Second)
   barTimer := time.NewTicker(60 * time.Second)
   for {
   	select {
   	case <-tickTimer.C:
   		fmt.Println("tick")
   	case <-barTimer.C:
   		fmt.Println("bar")
   	}
   }
}
```

# select 实现原理

Golang 实现 select 时，定义了一个数据结构表示每个 case 语句(含 defaut，default 实际上是一种特殊的 case)，select 执行过程可以类比成一个函数，函数输入 case 数组，输出选中的 case，然后程序流程转到选中的 case 块。

## case 数据结构

源码包 `src/runtime/select.go:scase` 定义了表示 case 语句的数据结构：

```go
type scase struct {
	c           *hchan         // chan
	kind        uint16
	elem        unsafe.Pointer // data element
}
```

scase.c 为当前 case 语句所操作的 channel 指针，这也说明了一个 case 语句只能操作一个 channel。scase.kind 表示该 case 的类型，分为读 channel、写 channel 和 default，三种类型分别由常量定义：

- caseRecv：case 语句中尝试读取 scase.c 中的数据；
- caseSend：case 语句中尝试向 scase.c 中写入数据；
- caseDefault： default 语句

scase.elem 表示缓冲区地址，跟据 scase.kind 不同，有不同的用途：

- scase.kind == caseRecv ： scase.elem 表示读出 channel 的数据存放地址；
- scase.kind == caseSend ： scase.elem 表示将要写入 channel 的数据存放地址；

## select 实现逻辑

源码包 src/runtime/select.go:selectgo() 定义了 select 选择 case 的函数：

```go
func selectgo(cas0 *scase, order0 *uint16, ncases int) (int, bool)
```

函数参数：

- cas0 为 scase 数组的首地址，selectgo()就是从这些 scase 中找出一个返回。
- order0 为一个两倍 cas0 数组长度的 buffer，保存 scase 随机序列 pollorder 和 scase 中 channel 地址序列 lockorder
  - pollorder：每次 selectgo 执行都会把 scase 序列打乱，以达到随机检测 case 的目的。
  - lockorder：所有 case 语句中 channel 序列，以达到去重防止对 channel 加锁时重复加锁的目的。
- ncases 表示 scase 数组的长度

函数返回值：

1. int： 选中 case 的编号，这个 case 编号跟代码一致
2. bool: 是否成功从 channle 中读取了数据，如果选中的 case 是从 channel 中读数据，则该返回值表示是否读取成功。

selectgo 实现伪代码如下：

```go
func selectgo(cas0 *scase, order0 *uint16, ncases int) (int, bool) {
    //1. 锁定scase语句中所有的channel
    //2. 按照随机顺序检测scase中的channel是否ready
    //   2.1 如果case可读，则读取channel中数据，解锁所有的channel，然后返回(case index, true)
    //   2.2 如果case可写，则将数据写入channel，解锁所有的channel，然后返回(case index, false)
    //   2.3 所有case都未ready，则解锁所有的channel，然后返回（default index, false）
    //3. 所有case都未ready，且没有default语句
    //   3.1 将当前协程加入到所有channel的等待队列
    //   3.2 当将协程转入阻塞，等待被唤醒
    //4. 唤醒后返回channel对应的case index
    //   4.1 如果是读操作，解锁所有的channel，然后返回(case index, true)
    //   4.2 如果是写操作，解锁所有的channel，然后返回(case index, false)
}

func selectgo(sel *hselect) int {
	// ...

	// case洗牌
	pollslice := slice{unsafe.Pointer(sel.pollorder), int(sel.ncase), int(sel.ncase)}
	pollorder := *(*[]uint16)(unsafe.Pointer(&pollslice))
	for i := 1; i < int(sel.ncase); i++ {
		//....
	}

	// 给case排序
	lockslice := slice{unsafe.Pointer(sel.lockorder), int(sel.ncase), int(sel.ncase)}
	lockorder := *(*[]uint16)(unsafe.Pointer(&lockslice))
	for i := 0; i < int(sel.ncase); i++ {
		// ...
	}
	for i := int(sel.ncase) - 1; i >= 0; i-- {
		// ...
	}

	// 加锁该select中所有的channel
	sellock(scases, lockorder)

	// 进入loop
loop:
	// ...
	// pass 1 - look for something already waiting
	// 按顺序遍历case来寻找可执行的case
	for i := 0; i < int(sel.ncase); i++ {
		//...
		switch cas.kind {
		case caseNil:
			continue
		case caseRecv:
			// ... goto xxx
		case caseSend:
			// ... goto xxx
		case caseDefault:
			dfli = casi
			dfl = cas
		}
	}

	// 没有找到可以执行的case，但有default条件，这个if里就会直接退出了。
	if dfl != nil {
		// ...
	}
	// ...

	// pass 2 - enqueue on all chans
	// chan入等待队列
	for _, casei := range lockorder {
		// ...
		switch cas.kind {
		case caseRecv:
			c.recvq.enqueue(sg)

		case caseSend:
			c.sendq.enqueue(sg)
		}
	}

	// wait for someone to wake us up
	// 等待被唤起,同时解锁channel(selparkcommit这里实现的)
	gp.param = nil
	gopark(selparkcommit, nil, "select", traceEvGoBlockSelect, 1)

	// 突然有故事发生，被唤醒，再次该select下全部channel加锁
	sellock(scases, lockorder)

	// pass 3 - dequeue from unsuccessful chans
	// 本轮最后一次循环操作，获取可执行case，其余全部出队列丢弃
	casi = -1
	cas = nil
	sglist = gp.waiting
	// Clear all elem before unlinking from gp.waiting.
	for sg1 := gp.waiting; sg1 != nil; sg1 = sg1.waitlink {
		sg1.isSelect = false
		sg1.elem = nil
		sg1.c = nil
	}
	gp.waiting = nil

	for _, casei := range lockorder {
		// ...
		if sg == sglist {
			// sg has already been dequeued by the G that woke us up.
			casi = int(casei)
			cas = k
		} else {
			c = k.c
			if k.kind == caseSend {
				c.sendq.dequeueSudoG(sglist)
			} else {
				c.recvq.dequeueSudoG(sglist)
			}
		}
		// ...
	}

	// 没有的话，再走一次loop
	if cas == nil {
		goto loop
	}
	// ...
bufrecv:
	// can receive from buffer
bufsend:
	// ...
recv:
	// ...
rclose:
	// ...
send:
	// ...
retc:
	// ...
sclose:
	// send on closed channel
}
```

特别说明：对于读 channel 的 case 来说，如 `case elem, ok := <-chan1:`, 如果 channel 有可能被其他协程关闭的情况下，一定要检测读取是否成功，因为 close 的 channel 也有可能返回，此时 ok == false。

![selectgo 流程分析](https://s1.ax1x.com/2020/06/07/tRDoxU.md.png)
