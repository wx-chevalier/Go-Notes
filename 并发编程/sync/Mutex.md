# Mutex

Mutex 代表 Mutual Exclusion(互斥)。互斥提供了一种并发安全的方式来表示对共享资源访问的独占。下面是一个简单的两个 goroutine，它们试图增加和减少一个公共值，并使用 Mutex 来同步访问：

```go
var count int
var lock sync.Mutex

increment := func() {
	lock.Lock() // 1
	defer lock.Unlock() // 2
	count++
	fmt.Printf("Incrementing: %d\n", count)
}

decrement := func() {
	lock.Lock() // 1
	defer lock.Unlock() // 2
	count--
	fmt.Printf("Decrementing: %d\n", count)
}

// Increment
var arithmetic sync.WaitGroup
for i := 0; i <= 5; i++ {
	arithmetic.Add(1)
	go func() {
		defer arithmetic.Done()
		increment()
	}()

}

// Decrement
for i := 0; i <= 5; i++ {
	arithmetic.Add(1)
	go func() {
		defer arithmetic.Done()
		decrement()
	}()
}

arithmetic.Wait()
fmt.Println("Arithmetic complete.")
```
