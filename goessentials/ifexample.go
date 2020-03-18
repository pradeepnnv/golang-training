package main

import (
	"fmt"
)

func main() {
	x := 5
	if x >= 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is less than 5")
	}

	if x > 3 && x < 5 {
		fmt.Println("X is between 3 & 5")
	} else {
		fmt.Println("X is not between 3 & 5")
	}
	if x < 4 || x > 5 {
		fmt.Println("x is between 3 & 5")
	}

	x = 9
	y := 10
	if z := x / y; z < 5 {
		fmt.Println("Fraction is less than 5")
	}
}
