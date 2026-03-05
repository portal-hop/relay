package main

import (
	"fmt"
	"io"
	"net"
)

// Takes packets from port A
// Sends to port B

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server running on :8081")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go handleConnection(conn) // One goroutine per connection
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	server, err := net.Dial("tcp", "localhost:25565")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer server.Close()

	go io.Copy(server, conn)
	io.Copy(conn, server)
}
