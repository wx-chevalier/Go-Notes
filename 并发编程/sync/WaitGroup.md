# WaitGroup

如果你不关心并发操作的结果，或者有其他方式收集结果，那么 WaitGroup 是等待一组并发操作完成的好方法。如果这两个条件都不成立，我建议你改用 channel 和 select 语句。

```go
var wg sync.WaitGroup

wg.Add(1) //1
go func() {
	defer wg.Done() //2
	fmt.Println("1st goroutine sleeping...")
	time.Sleep(1)
}()

wg.Add(1) //1
go func() {
	defer wg.Done() //2
	fmt.Println("2nd goroutine sleeping...")
	time.Sleep(2)
}()

wg.Wait() //3
fmt.Println("All goroutines complete.")
```

可以把 WaitGroup 视作一个安全的并发计数器：调用 Add 增加计数，调用 Done 减少计数。调用 Wait 会阻塞并等待至计数器归零。注意，Add 的调用是在 goroutines 之外完成的。如果没有这样做，我们会引入一个数据竞争条件，因为我们没有对 goroutine 做任何调度顺序上的保证; 我们可能在任何一个 goroutines 开始前触发 Wait 调用。如果 Add 的调用被放置在 goroutines 的闭包中，对 Wait 的调用可能完全没有阻塞地返回，因为 Add 没有被执行。

通常情况下，尽可能与要跟踪的 goroutine 就近且成对的调用 Add，但有时候会一次性调用 Add 来跟踪一组 goroutine。我通常会做这样的循环：

```go
hello := func(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Hello from %v!\n", id)
}

const numGreeters = 5
var wg sync.WaitGroup
wg.Add(numGreeters)
for i := 0; i < numGreeters; i++ {
	go hello(&wg, i+1)
}
wg.Wait()
```
