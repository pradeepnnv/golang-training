//This prints fizz if a number is divisible by 3 & buzz if it's divisible by 10

package main

import (
	"fmt"
)

func main() {
	i := 0
	for i = 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
