package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Author  string
}

func main() {
	file, _ := os.Open("conf.json")

	defer file.Close()

	// A decoder is used to decode a stream of data
	decoder := json.NewDecoder(file)
	conf := configuration{}

	var err = decoder.Decode(&conf)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conf.Author)

}
