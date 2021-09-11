// Setting up a log file and send log messages to that
package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./logger.txt")
	defer logfile.Close()

	logger := log.New(logfile, "logger", log.LstdFlags|log.LstdFlags)

	logger.Println("Hello world!")
	logger.Fatalln("Oh snap! Error happened")
}
