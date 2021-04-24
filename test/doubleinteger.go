package mytest

import "fmt"

func main() {
	fmt.Println(doubleTheInteger(4))
}

func doubleTheInteger(i int) int {
	return i * 2
}
