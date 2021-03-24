package mytest

import "fmt"

func main() {
	fmt.Println(doubleTheInteger(2))
}

func doubleTheInteger(i int) int {
	return i * 2
}
