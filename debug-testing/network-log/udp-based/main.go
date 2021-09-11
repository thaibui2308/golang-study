package main

import (
	"log"
	"net"
	"time"
)

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:3000", timeout)
	if err != nil {
		panic("Fail to connect to UDP server")
	}
	defer conn.Close()

	format := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", format)

	logger.Println("A regular message")
	logger.Panicln("A panic message")
}
