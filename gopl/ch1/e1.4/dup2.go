//Program to print all filenames which have duplicate lines

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Function to identify all files with duplicate lines
//File names are provided as command line arguments
func main() {

	if len(os.Args) <= 1 {
		fmt.Println("No filenames were provided as arguments")
	}
	for _, filename := range os.Args[1:] {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2:%s", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		hasDups := false
		for _, n := range counts {
			if n > 1 {
				hasDups = true
			}
		}
		if hasDups {
			fmt.Printf("%s has duplicate lines\n", filename)
		}
	}
}
