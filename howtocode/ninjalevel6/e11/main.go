package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

// factorial using loops

// func factorial(x int) int {
// 	f := x
// 	for i := 1; i < x; i++ {
// 		f = f * (x - i)
// 	}
// 	return f
// }

// factorial using recursion

func factorial(x int) int {
	if x < 1 {
		return 1
	}
	return x * factorial(x-1)
}
