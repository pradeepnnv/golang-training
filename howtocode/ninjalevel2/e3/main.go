package main

import "fmt"

const x = 42
const y int = 94

const (
	a        = 42
	b string = "const string"
)

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a, b)
}
