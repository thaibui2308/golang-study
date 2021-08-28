package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 2)

	for i := 1; i < 10; i++ {
		go worker(i, lock)
	}

	time.Sleep(7 * time.Second)
}

func worker(i int, lock chan bool) {
	fmt.Printf("%d wants the lock \n", i)
	lock <- true
	fmt.Printf("%d has the lock\n", i)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", i)
	// release a lock by reading a value, which will open that one space on the buffer again
	<-lock
}
