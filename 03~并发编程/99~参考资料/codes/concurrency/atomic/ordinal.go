package atomic

import (
	"sync"
	"sync/atomic"
)

type Ordinal struct {
	ordinal uint64
	once    *sync.Once
}

/*
NewOrdinal is a
*/
func NewOrdinal() *Ordinal {
	return &Ordinal{once: &sync.Once{}}
}
func (o *Ordinal) Init(val uint64) {
	o.once.Do(func() {
		atomic.StoreUint64(&o.ordinal, val)
	})
}
func (o *Ordinal) GetOrdinal() uint64 {
	return atomic.LoadUint64(&o.ordinal)
}
func (o *Ordinal) Increment() {
	atomic.AddUint64(&o.ordinal, 1)
}
