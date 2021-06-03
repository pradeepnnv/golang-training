package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("conf.json")

	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(conf)
}
