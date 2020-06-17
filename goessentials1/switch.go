//Program to demonstrate switch example

package main

import "fmt"

func main() {
	x := 5
	switch {
	case x < 5:
		fmt.Println("X is less than 5")
	case (x >= 5 && x < 10):
		fmt.Println("X is greater than or equal to 5 and less than 10")
	case (x == 10):
		fmt.Println("X is equal to 10")
	case x > 10:
		fmt.Println("X is greater than 10")
	default:
		fmt.Println("Default")
	}
}
