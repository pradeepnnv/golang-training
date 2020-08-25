package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	fmt.Printf("Type of x is %T\n", x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of x is %T\n", x)
	y := []int{1, 2, 3, 4}
	fmt.Printf("Type of y is %T\n", y)
}
