package session

type SessionMgr interface {
	Init(options ...interface{}) error
	CreateSession() (session Session, err error)
	Get(sessionID string) (session Session, err error)
}
