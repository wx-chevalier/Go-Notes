package bench

import "sync/atomic"

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Add(amount int64) {
	atomic.AddInt64(&c.value, amount)
}

func (c *AtomicCounter) Read() int64 {
	var result int64
	result = atomic.LoadInt64(&c.value)
	return result
}
