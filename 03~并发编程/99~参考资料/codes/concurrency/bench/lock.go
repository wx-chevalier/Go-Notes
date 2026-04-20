package bench

import "sync"

type Counter struct {
	value int64
	mu    *sync.RWMutex
}

func (c *Counter) Add(amount int64) {
	c.mu.Lock()
	c.value += amount
	c.mu.Unlock()
}

func (c *Counter) Read() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
