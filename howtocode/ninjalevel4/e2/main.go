package main

import "fmt"

func main() {
	xs := []int{
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}
	for i := range xs {
		fmt.Println(i)
	}
	fmt.Printf("Type of xs is %T", xs)
}
