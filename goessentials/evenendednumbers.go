//This program identifies all even ended numbers resulted by multiplying two four digit numbers

package main

import (
	"fmt"
)

func main() {
	count := 0

	for i := 1001; i <= 9999; i++ {
		for j := 1001; j <= 9999; j++ {
			strNum := fmt.Sprintf("%d", i*j)
			if strNum[0] == strNum[len(strNum)-1] {
				count++
			}
		}

	}
	fmt.Println(count)
}
