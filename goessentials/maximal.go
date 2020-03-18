//maximal calculcates the biggest value in a slice of numbers
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 2, 3, 94, 2, 99, 98, 79, 4, 2}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	fmt.Println("Maximal number is :", max)
}
