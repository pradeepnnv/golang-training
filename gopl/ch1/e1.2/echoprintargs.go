//Program to print arguments

package main

import (
	"fmt"
	"os"
)

//Main method
func main() {
	for i, v := range os.Args {
		fmt.Printf("%d=====>%s\n", i, v)
	}
}
