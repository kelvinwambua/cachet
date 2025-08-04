package server

import (
	"cachet/internal/store"
	"fmt"
	"net"
)

type Server struct {
	addr  string
	store store.Store
}

func New(addr string, store store.Store) *Server {
	return &Server{
		addr:  addr,
		store: store,
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", s.addr, err)
	}
	defer listener.Close()

	fmt.Printf("Cachet server started on %s\n", s.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go s.handleConnection(conn)
	}
}
