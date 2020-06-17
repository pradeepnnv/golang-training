//Program to determine mean of two numbers
package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	x = 5
	y = 6
	var mean float64
	mean = (5 + 6) / 2.0
	fmt.Printf("Mean of %d & %d is %f", x, y, mean)
	fmt.Println(mean)
}
