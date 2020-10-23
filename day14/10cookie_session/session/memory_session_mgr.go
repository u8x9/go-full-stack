package session

import (
	"fmt"
	"sync"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}

func (msm *MemorySessionMgr) Init(options ...interface{}) error {
	return nil
}
func (msm *MemorySessionMgr) CreateSession() (session Session, err error) {
	msm.rwLock.Lock()
	defer msm.rwLock.Unlock()
	sessionID := NewSessionID()
	session = NewMemorySession(sessionID)
	msm.sessionMap[sessionID] = session
	return
}
func (msm *MemorySessionMgr) Get(sessionID string) (session Session, err error) {
	var ok bool
	msm.rwLock.RLock()
	defer msm.rwLock.RUnlock()
	if session, ok = msm.sessionMap[sessionID]; !ok {
		return nil, fmt.Errorf("the session %q is not exists", sessionID)
	}
	return session, nil
}
