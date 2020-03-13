//Echo prints command line arguments including Args0
//Excercise 1.1
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0:])
}
