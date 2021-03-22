package mux

import (
	"slowhttp/handler"
	"slowhttp/req"
)

type Mux struct {
	matcher map[string]handler.Handler
}

func NewMux() *Mux {
	return &Mux{matcher: make(map[string]handler.Handler)}
}

func (m *Mux) Match(req *req.Request) handler.Handler {
	return m.matcher[req.Path]
}

func (m *Mux) Reg(path string, h handler.Handler) {
	m.matcher[path] = h
}
