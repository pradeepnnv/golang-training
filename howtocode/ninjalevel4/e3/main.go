package main

import "fmt"

func main() {
	xs := []int{
		42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}

	fmt.Println(xs[0:5])
	fmt.Println(xs[5:])
	fmt.Println(xs[2:7])
	fmt.Println(xs[1:6])
}
