package main

import (
	"flag"
	"fmt"
	"os"
)

// First way to create a flag
var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool
var exit bool

func init() {
	// Second way to create a flag. Long and short flags
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")

	// Exit flag
	flag.BoolVar(&exit, "exit", false, "Use to escape the current executing process")
}

func main() {
	flag.Parse()

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}

	if exit == true {
		os.Exit(0)
	}

	// Custom flag help text, accecpting a callback function that iterates over each flags executing the callback function
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
}
