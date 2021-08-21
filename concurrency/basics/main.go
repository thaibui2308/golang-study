package main

import (
	"fmt"
	"io"
	"runtime"
)

func main() {

	fmt.Println("Outside a go routine")

	go func() {
		fmt.Println("Inside a go routine pt.1")
	}()
	go func() {
		fmt.Println("Inside a go routine pt.2")
	}()
	runtime.Gosched()
	fmt.Println("Outside again.")

	// Each line u type in will be displayed by the shell
	// go echo(os.Stdin, os.Stdout)
	// time.Sleep(40 * time.Second)
	// fmt.Println("Time out!")
	// os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	// Copies data to an os.Writer from an os.Reader
	io.Copy(out, in)
}
