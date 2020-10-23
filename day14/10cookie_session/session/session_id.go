package session

import "github.com/rs/xid"

func NewSessionID() string {
	return xid.New().String()
}
