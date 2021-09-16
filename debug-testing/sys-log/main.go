package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	priority := 2 | 3
	flags := log.Ldate | log.Ltime | log.Lshortfile

	logger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		fmt.Printf("Can't attach to syslog: %s", err)
		return
	}
	logger.Println("Test log message")
}
