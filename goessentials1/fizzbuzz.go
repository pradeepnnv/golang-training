// FizzBuzz
// This program prints fizz if the number is divisible by 3 and buzz if it's divisible by 5. It'll print fizzbuzz if it's divisible by 3 & 5.

package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		switch {
		case (i%3 == 0 && i%5 == 0):
			fmt.Println("fizzbuzz")
		case (i%3 == 0):
			fmt.Println("fizz")
		case (i%5 == 0):
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
