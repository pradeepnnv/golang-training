//Program which has a function to double the value of an integer at a given index in a slice

package main

import "fmt"

func main() {
	i := 10
	fmt.Println("Before doubling", i)
	double(i)
	fmt.Println("After doubling:", i)

	doublePtr(&i)
	fmt.Println("After doubling through pointers", i)
}

func double(i int) {
	i = i * 2
}

func doublePtr(i *int) {
	*i = *i * 2
}
