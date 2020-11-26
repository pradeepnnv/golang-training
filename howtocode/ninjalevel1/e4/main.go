package main

import "fmt"

func main() {
	type mytype int
	var x mytype
	x = 42
	fmt.Printf("Variable x is of type %T and value is %d", x, x)
}
