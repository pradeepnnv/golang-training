//This program prints maximal value in a slice of numbers

package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	x := nums[0]

	for _, v := range nums {

		if v > x {
			x = v
		}
	}
	fmt.Println(x)
}
