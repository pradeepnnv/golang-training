package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println(add(1, 2))
	fmt.Println(divmod(1, 2))
}

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}
