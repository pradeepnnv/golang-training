//Program to find maximal value in a slice of randomly chosen numbers

package main

import "fmt"

func main() {
	maxnum := 0
	nums := []int{3, 39, 4, 23, 244, 3, 9, 6}
	for _, v := range nums {
		if v > maxnum {
			maxnum = v
		}
	}
	fmt.Println(maxnum)
}
