package server

import (
	"cachet/internal/store"
	"fmt"
	"net"
	"strconv"
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
func (s *Server) handleIncr(args []string) string {
	if len(args) != 1 {
		return "ERROR: INCR requires exactly 1 argument"
	}

	key := args[0]
	value, exists := s.store.Get(key)

	var num int64 = 0
	if exists {
		var err error
		num, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return "ERROR: value is not an integer"
		}
	}

	num++
	s.store.Set(key, strconv.FormatInt(num, 10))
	return strconv.FormatInt(num, 10)
}

func (s *Server) handleDecr(args []string) string {
	if len(args) != 1 {
		return "ERROR: DECR requires exactly 1 argument"
	}

	key := args[0]
	value, exists := s.store.Get(key)

	var num int64 = 0
	if exists {
		var err error
		num, err = strconv.ParseInt(value, 10, 64)
		if err != nil {
			return "ERROR: value is not an integer"
		}
	}

	num--
	s.store.Set(key, strconv.FormatInt(num, 10))
	return strconv.FormatInt(num, 10)
}

func (s *Server) handleAppend(args []string) string {
	if len(args) != 2 {
		return "ERROR: APPEND requires exactly 2 arguments"
	}

	key, appendValue := args[0], args[1]
	existingValue, exists := s.store.Get(key)

	var newValue string
	if exists {
		newValue = existingValue + appendValue
	} else {
		newValue = appendValue
	}

	s.store.Set(key, newValue)
	return strconv.Itoa(len(newValue))
}

func (s *Server) handleStrlen(args []string) string {
	if len(args) != 1 {
		return "ERROR: STRLEN requires exactly 1 argument"
	}

	key := args[0]
	value, exists := s.store.Get(key)
	if !exists {
		return "0"
	}
	return strconv.Itoa(len(value))
}
