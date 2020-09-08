package main

import "fmt"

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5))
	fmt.Println(boo([]int{3, 4, 5, 6}))
}

func foo(vals ...int) int {
	sum := 0
	for _, i := range vals {
		sum += i
	}
	return sum
}

func boo(vals []int) int {
	sum := 0
	for _, i := range vals {
		sum += i
	}
	return sum
}
