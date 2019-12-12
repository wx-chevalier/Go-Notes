# 线程协作

Go 遵循称为 fork-join 模型的并发模型.fork 这个词指的是在程序中的任何一点，它都可以将一个子执行的分支分离出来，以便与其父代同时运行。join 这个词指的是这样一个事实，即在将来的某个时候，这些并发的执行分支将重新组合在一起，子分支重新加入的地方称为连接点。

![fork-join 模型](https://s2.ax1x.com/2019/12/11/QsvaTS.png)

所谓的连接点就是来同步 main goroutine 和 sayHello goroutine。

```go
var wg sync.WaitGroup
sayHello := func() {
	defer wg.Done()
	fmt.Println("hello")
}
wg.Add(1)
go sayHello()
wg.Wait() //1
```
