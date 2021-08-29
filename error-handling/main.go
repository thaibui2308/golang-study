package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// In the case of two or more errors and the application needs to specify which error it encounters,
// create a package-scoped variable that encapsulate the possible errors
var ErrTimeout = errors.New("the request timed out")
var ErrRejected = errors.New("the request was rejected")

var random = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	// Handling one error effectively
	args := os.Args[1:]

	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Concatenated string: %s\n", result)
	}

	// Dealing with two errors
	response, err := SendRequest("Thai Bui")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying...")
		response, err = SendRequest("Thai Bui")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

// Handling one error
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no strings supplied")
	}

	return strings.Join(parts, " "), nil
}

func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
