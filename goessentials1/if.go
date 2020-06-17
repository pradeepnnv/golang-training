//If Example in Go

package main

import (
	"fmt"
)

func main() {

	x := 25

	if x < 5 {
		fmt.Println("X is less than 5")
	} else if x >= 5 && x < 10 {
		fmt.Println("X is between 5 & 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is greater than 10")
	}

	if y := (x % 5); y < 5 {
		fmt.Println("Y is less than 5")
	}
}
