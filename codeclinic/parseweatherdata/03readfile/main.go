package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("main.go")

	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Println(f)
}
