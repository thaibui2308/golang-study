package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var i int = -1

	var file string

	for i, file = range os.Args[1:] {
		// Tell the waitgroup u hafta wait for one more compress operation
		wg.Add(1)
		go func(filename string) {
			compress(filename)
			// Notify the waitgroup that the file in line is done
			wg.Done()
		}(file)
	}
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// Compress data and then writes it to the underlying file
	gzout := gzip.NewWriter(out)
	// Copy the data from the original file to the gzip.Writer and it will be written to the .gz file
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err

}
