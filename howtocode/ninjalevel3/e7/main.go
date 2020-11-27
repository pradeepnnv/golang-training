package main

import "fmt"

func main() {
	x := 42
	if x < 40 {
		fmt.Println("X is less than 40")
	} else if x > 45 {
		fmt.Println("X is greater than 40")
	}
}
