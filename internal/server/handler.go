package server

import (
    "bufio"
    "fmt"
    "net"
    "strconv"
    "strings"
)

func (s *Server) handleConnection(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)

    conn.Write([]byte("Welcome to Cachet!\n"))

    for {
        conn.Write([]byte("> "))

        if !scanner.Scan() {
            break
        }

        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }

        response := s.processCommand(line)
        conn.Write([]byte(response + "\n"))
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Connection error: %v\n", err)
    }
}

func (s *Server) processCommand(line string) string {
    parts := strings.Fields(line)
    if len(parts) == 0 {
        return "ERROR: empty command"
    }

    command := strings.ToUpper(parts[0])
    args := parts[1:]

    switch command {
    case "GET":
        return s.handleGet(args)
    case "SET":
        return s.handleSet(args)
    case "DEL", "DELETE":
        return s.handleDelete(args)
    case "EXISTS":
        return s.handleExists(args)
    case "KEYS":
        return s.handleKeys(args)
    case "SIZE":
        return s.handleSize(args)
    case "CLEAR":
        return s.handleClear(args)
    case "PING":
        return "PONG"
    case "INCR":
        return s.handleIncr(args)
    case "DECR":
        return s.handleDecr(args)
    case "APPEND":
        return s.handleAppend(args)
    case "STRLEN":
        return s.handleStrlen(args)
    case "EXIT", "QUIT":
        return "Goodbye!"
    default:
        return fmt.Sprintf("ERROR: unknown command '%s'", command)
    }
}

func (s *Server) handleGet(args []string) string {
    if len(args) != 1 {
        return "ERROR: GET requires exactly 1 argument"
    }

    key := args[0]
    value, exists := s.store.Get(key)
    if !exists {
        return "(nil)"
    }
    return value
}

func (s *Server) handleSet(args []string) string {
    if len(args) != 2 {
        return "ERROR: SET requires exactly 2 arguments"
    }

    key, value := args[0], args[1]
    s.store.Set(key, value)
    return "OK"
}

func (s *Server) handleDelete(args []string) string {
    if len(args) != 1 {
        return "ERROR: DELETE requires exactly 1 argument"
    }

    key := args[0]
    deleted := s.store.Delete(key)
    if deleted {
        return "1"
    }
    return "0"
}

func (s *Server) handleExists(args []string) string {
    if len(args) != 1 {
        return "ERROR: EXISTS requires exactly 1 argument"
    }

    key := args[0]
    exists := s.store.Exists(key)
    if exists {
        return "1"
    }
    return "0"
}

func (s *Server) handleKeys(args []string) string {
    keys := s.store.Keys()
    if len(keys) == 0 {
        return "(empty list)"
    }

    result := strings.Join(keys, ", ")
    return fmt.Sprintf("[%s]", result)
}

func (s *Server) handleSize(args []string) string {
    size := s.store.Size()
    return strconv.Itoa(size)
}

func (s *Server) handleClear(args []string) string {
    s.store.Clear()
    return "OK"
}
