package main

import "fmt"

func main() {
	xs := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(xs)
	for ix, x := range xs {
		fmt.Println("This is :", ix)
		for _, v := range x {
			fmt.Println(v)
		}
	}
}
