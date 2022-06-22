package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for _, filename := range os.Args[1:] {
		counts := make(map[string]int)
		f, err := os.Open(filename)
		defer f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v", err)
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}

		for _, n := range counts {
			if n > 1 {
				fmt.Printf("File %s has duplicates\n", filename)
				break
			}
		}
	}
}
