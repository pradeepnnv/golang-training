package main

import "fmt"

var x int
var y string

func main() {
	x = 42
	y = "hello world"

	fmt.Printf("Var x is of type %T and value %d\n", x, x)
	fmt.Printf("Variable y is of type %T and value is %s\n", y, y)
}
