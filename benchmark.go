package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	const addr = "localhost:6380"
	const numRequests = 1000000000000000000

	start := time.Now()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for i := 0; i < numRequests; i++ {
		cmd := fmt.Sprintf("SET key%d value%d\n", i, i)
		_, err := conn.Write([]byte(cmd))
		if err != nil {
			panic(err)
		}
		_, _ = reader.ReadString('\n')
	}

	elapsed := time.Since(start)
	fmt.Printf("Completed %d SET requests in %s\n", numRequests, elapsed)
	fmt.Printf("Average time per request: %s\n", elapsed/time.Duration(numRequests))
}
