package main

import "fmt"

func main() {
	a := 42
	b := 32

	g := a == b
	h := a <= b
	i := a >= b
	j := a != b
	k := a < b
	l := a > b

	fmt.Println(g, h, i, j, k, l)

}
