package server

import (
	"cachet/internal/db"
	"fmt"
	"net"
)

type Server struct {
	addr  string
	store *db.Store
}

func New(addr string, store *db.Store) *Server {
	return &Server{addr: addr, store: store}
}

func (s *Server) ListenAndServe() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	fmt.Println("Server started on", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn, s.store)
	}
}
