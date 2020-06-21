//Program to print all arguments along with the program name

package main

import (
	"fmt"
	"os"
	"time"
)

//Main Function
func main() {
	start := time.Now()

	//Print directly
	// fmt.Println(os.Args)

	//Print using strings.Join
	// fmt.Println(strings.Join(os.Args[1:], " "))

	//Print using array
	// for _, s := range os.Args {
	// 	fmt.Println(s)
	// }

	//Print using string concatenation
	s, sep := "", " "
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}
