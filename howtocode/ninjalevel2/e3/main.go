package main

import "fmt"

func main() {
	const x = 22
	const y string = "hello"

	fmt.Printf("Type of x is %T and value is %d\n", x, x)
	fmt.Printf("Type of Y is %T and value is %s\n", y, y)
}
