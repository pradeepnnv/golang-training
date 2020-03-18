//evenended prints all even ended numbers between the specified range. A number is considered even if first and last digits are same.

package main

import (
	"fmt"
)

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)
	fmt.Printf(s)
	counter := 0
	for i := 1000; i < 9999; i++ {
		for j := i; j < 9999; j++ {
			s := fmt.Sprintf("%d", i*j)
			if s[0] == s[len(s)-1] {
				counter++
			}
		}
	}
	fmt.Println("Count of four digit even ended numbers is :", counter)
}
