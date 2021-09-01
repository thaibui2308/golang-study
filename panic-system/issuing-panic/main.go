package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		// recover function retrieves the data from the panic
		if err := recover(); err != nil {
			// return a value (interface{}) if a panic has been raised, but in all other cases it returns nil
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	yikes()
}

func yikes() {
	panic(errors.New("unidentified problem"))
}
