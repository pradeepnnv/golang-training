// Multiples of 3 & 5
// Problem statement: https://projecteuler.net/problem=1

package main

import "fmt"

// Function to identify sum of all multiples of 3 & 5 between 1 & 1000
func main() {
	sumValue := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sumValue += i
		}
	}
	fmt.Println("Answer is :", sumValue)
}
