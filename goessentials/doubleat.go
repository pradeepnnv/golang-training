package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
	fmt.Println(values)
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4, 5}
	doubleAt(values, 4)
	fmt.Println(values)
	val := 2
	double(val)

	doublePtr(&values[2])
	fmt.Println(val)
	fmt.Println(values)
}
