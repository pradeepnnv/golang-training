package main

import "fmt"

func main() {
	var xs [5]int
	xs[0] = 1
	xs[1] = 2
	fmt.Println(xs)
	fmt.Printf("Type of xs is %T", xs)
}
