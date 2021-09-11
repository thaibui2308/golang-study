package main

import (
	"log"
	"net"
)

func main() {
	// Dial connects to a server that already initialized beforehand
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		panic("Fail to connect to localhost:3000")
	}

	defer conn.Close()

	logger := log.New(conn, "example", log.Ldate|log.Lshortfile)
	// Send log message to the network connection.
	logger.Printf("Connect to localhost:3000")
	// Do not use Fatal as it will abruptly disconnect to the server and you wont have a chance to
	// close the connection with the above defer function.
	logger.Panic("Oh snap! I'm panicking!")
}
