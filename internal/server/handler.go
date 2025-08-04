package server

import (
    "bufio"
    "cachet/internal/db"
    "net"
    "strings"
)

func handleConnection(conn net.Conn, store *db.Store) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    for {
        conn.Write([]byte("> "))
        line, err := reader.ReadString('\n')
        if err != nil {
            return
        }

        line = strings.TrimSpace(line)
        parts := strings.SplitN(line, " ", 3)

        if len(parts) == 0 {
            continue
        }

        switch strings.ToUpper(parts[0]) {
        case "SET":
            if len(parts) != 3 {
                conn.Write([]byte("Usage: SET key value\n"))
                continue
            }
            store.Set(parts[1], parts[2])
            conn.Write([]byte("OK\n"))
        case "GET":
            if len(parts) != 2 {
                conn.Write([]byte("Usage: GET key\n"))
                continue
            }
            val, ok := store.Get(parts[1])
            if !ok {
                conn.Write([]byte("(nil)\n"))
            } else {
                conn.Write([]byte(val + "\n"))
            }
        case "EXIT":
            conn.Write([]byte("Goodbye!\n"))
            return
        default:
            conn.Write([]byte("Unknown command\n"))
        }
    }
}
