package main

import "fmt"

type mynewtype int

var x mynewtype
var y int

func main() {
	x := 42
	y := 48
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	x = y
	fmt.Println(x, y)
}
