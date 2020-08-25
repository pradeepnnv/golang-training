package main

import "fmt"

func main() {
	xs := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypeynny", "Hiyaa, James"},
	}
	for _, i := range xs {
		fmt.Println("\t")
		for _, j := range i {
			fmt.Println("\t", j)
		}
	}
}
