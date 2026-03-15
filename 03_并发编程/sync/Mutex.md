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

很多时候我们希望对某个对象屏蔽并发访问，可以将 Mutex 添加伪结构体的某个属性：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	Name   string
	Locker *sync.Mutex
}

func (u *User) SetName(wati *sync.WaitGroup, name string) {
	defer func() {
		fmt.Println("Unlock set name:", name)
		u.Locker.Unlock()
		wati.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock set name:", name)
	time.Sleep(1 * time.Second)
	u.Name = name
}

func (u *User) GetName(wati *sync.WaitGroup) {
	defer func() {
		fmt.Println("Unlock get name:", u.Name)
		u.Locker.Unlock()
		wati.Done()
	}()

	u.Locker.Lock()
	fmt.Println("Lock get name:", u.Name)
	time.Sleep(1 * time.Second)
}

func main() {
	user := User{}
	user.Locker = new(sync.Mutex)
	wait := &sync.WaitGroup{}
	names := []string{"a", "b", "c"}
	for _, name := range names {
		wait.Add(2)
		go user.SetName(wait, name)
		go user.GetName(wait)
	}

	wait.Wait()
}
```
