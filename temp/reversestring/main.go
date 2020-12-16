package main

import "fmt"

func main() {
	x := []string{"H", "E", "L", "L", "O"}
	fmt.Println(x)
	y := []string{""}
	for i := 0; i < len(x); i++ {
		y = append(y, x[len(x)-i-1])
	}
	fmt.Println(y)
}
