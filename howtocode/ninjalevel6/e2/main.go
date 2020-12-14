package main

import "fmt"

func main() {

	xs := []int{
		1, 2, 3, 4, 5, 6,
	}
	fmt.Println(findSum(xs...))
	fmt.Println(findSumSlice(xs))
}

func findSum(x ...int) int {
	var sum int
	for i := range x {
		sum += i
	}
	return sum
}

func findSumSlice(x []int) int {
	var sum int
	for i := range x {
		sum += i
	}
	return sum
}
