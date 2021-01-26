package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("logdetails.txt")
	if err != nil {
		log.Fatal("File could not be opened", err)
	}
	input := bufio.NewScanner(f)

	for input.Scan() {
		fmt.Println(input.Text())
	}
}
