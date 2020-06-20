// Program to print all command line arguments including the program name

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Println(i, ":", v)
	}
}
