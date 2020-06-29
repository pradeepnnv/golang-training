package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) == 0 {
		printFile(os.Stdin)
	} else {
		for _, fileName := range os.Args[1:] {
			inputFileHandler, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			printFile(inputFileHandler)
		}
	}
}

func printFile(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}
