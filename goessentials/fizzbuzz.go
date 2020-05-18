// Print fizz for numbers divisible by 3, buzz for numbers divisible by 5 and fizzbuzz if they are divisible by 3 & 5

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizz buzz")
		} else if i%3 == 0 {
			fmt.Printf("fizz")
		} else if i%5 == 0 {
			fmt.Printf("buzz")
		} else {
			fmt.Printf("%v", i)
		}
		fmt.Println()
	}
}
