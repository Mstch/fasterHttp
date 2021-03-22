package session

import "sync"

//session:conn 1:n
type Session struct {
	Id   string
	Meta *sync.Map
}
