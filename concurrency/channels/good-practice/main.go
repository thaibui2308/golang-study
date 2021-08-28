package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(5 * time.Second)

	go send(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(time.Millisecond * 6)
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(time.Millisecond * 500)
		}
	}
}
