package main

import "fmt"

type myint int

func main() {
	var x myint
	var y int
	x = 42
	y = int(x)

	fmt.Printf("Type of y is %T and value is %d\n", y, y)
	fmt.Printf("Type of x is %T and value is %d\n", x, x)
}
