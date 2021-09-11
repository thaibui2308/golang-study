package main

import "log"

func main() {
	// Write a message to os.Stder
	log.Println("No error occurred")
	// Write a message to os.Stder and then exits with an error code
	log.Fatalln("Error occurred")
	log.Println("Never executed")
}
