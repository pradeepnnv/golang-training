package main

import "fmt"

func main() {
	i := 94
	switch {
	case i > 24:
		fmt.Println("I is greater than 24")
	case i < 24:
		fmt.Println("I is less than 24")
	default:
		fmt.Println("I is equal to 24")
	}
}
