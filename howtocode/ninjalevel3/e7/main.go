package main

import "fmt"

func main() {
	a := 34
	b := 34
	if a > b {
		fmt.Println("A is greater than b")
	} else if a < b {
		fmt.Println("A is less than b")
	} else {
		fmt.Println("A is equal to b")
	}
}
