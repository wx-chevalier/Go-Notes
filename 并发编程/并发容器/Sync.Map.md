# Sync.Map

在 Go1.9 之前，Go 自带的 Map 不是并发安全的，因此我们需要自己再封装一层，给 Map 加上把读写锁,例如:

```go
type MapWithLock struct {
    sync.RWMutex
    M map[string]Kline
}
```
