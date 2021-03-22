package server

import (
	"net"
	"slowhttp/handler"
	"slowhttp/mux"
	"slowhttp/server/conn"
	"sync"
)

type Server struct {
	l        net.Listener
	m        *mux.Mux
	SessionK string
	Sessions *sync.Map
}

func NewServer(bind string) (*Server, error) {
	l, err := net.Listen("tcp", bind)
	if err != nil {
		return nil, err
	}
	return &Server{l: l, m: mux.NewMux()}, nil
}

func (s *Server) Serve() error {
	for {
		c, err := s.l.Accept()
		if err != nil {
			return err
		}
		ss := conn.CreateHttpConn(c, s.m)
		go ss.Serve()
		println(c, "new conn created")
	}
}

func (s *Server) RegHandler(path string, handler handler.Handler) {
	s.m.Reg(path, handler)
}

func (s *Server) Shutdown() {
	s.l.Close()
}
