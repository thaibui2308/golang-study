package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	// Start a new server on port 1026
	listener, err := net.Listen("tcp", ":3030")
	if err != nil {
		fmt.Println("Failed to open port on 3030")
		return
	}

	// Listens for new clients connections and handles any connection errors
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection")
			continue
		}

		// When a connection is accepted, passes it to the handle function
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s", err)
			conn.Close()
		}
	}()

	reader := bufio.NewReader(conn)

	// Read a line of data from the connection
	data, err := reader.ReadBytes('\n')
	// if fail, print an error and closes the connection
	if err != nil {
		fmt.Println("Failed to read from socket")
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	// write the data back out to the socket, echoing it to the client; then closes the connection.
	defer func() {
		conn.Close()
	}()
	conn.Write(data)
}
