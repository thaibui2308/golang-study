package main

import (
	"fmt"
	"os"
	"time"
)

var dataArr []int

func main() {
	// Wait for the specified second to elapse and then send the signal to the received channel
	done := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out!")
			os.Exit(0)
		}

	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
