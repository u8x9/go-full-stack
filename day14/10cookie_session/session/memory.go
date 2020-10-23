package session

import (
	"fmt"
	"sync"
)

type MemorySession struct {
	id     string
	data   SessionData
	rwLock sync.RWMutex
}

func NewMemorySession(id string) Session {
	return &MemorySession{
		id:   id,
		data: make(SessionData, 16),
	}
}

func (ms *MemorySession) Set(key string, value interface{}) error {
	ms.rwLock.Lock()
	ms.data[key] = value
	ms.rwLock.Unlock()
	return nil
}
func (ms *MemorySession) Get(key string) (value interface{}, err error) {
	var ok bool
	ms.rwLock.RLock()
	value, ok = ms.data[key]
	ms.rwLock.RUnlock()
	if ok {
		return value, nil
	}
	return nil, fmt.Errorf("The session %q not exists", key)
}
func (ms *MemorySession) Del(key string) error {
	ms.rwLock.Lock()
	delete(ms.data, key)
	ms.rwLock.Unlock()
	return nil
}
func (ms *MemorySession) Save() error {
	return fmt.Errorf("the MemorySession unimplemented 'Save' method")
}
