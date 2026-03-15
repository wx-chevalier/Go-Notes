package atomic

import (
	"errors"
	"sync"
)

type SafeMap struct {
	m  map[string]string
	mu *sync.RWMutex
}

func NewSafeMap() SafeMap {
	return SafeMap{m: make(map[string]string), mu: &sync.RWMutex{}}
}

func (t *SafeMap) Set(key, value string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.m[key] = value
}

func (t *SafeMap) Get(key string) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if v, ok := t.m[key]; ok {
		return v, nil
	}

	return "", errors.New("key not found")
}
