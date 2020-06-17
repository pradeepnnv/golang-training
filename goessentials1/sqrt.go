//Program to calculate square root of an integer. Returns zero with an error if it's a negative value.

package main

import (
	"fmt"
	"math"
)

func main() {
	s, err := squareRoot(936)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
	s, err = squareRoot(-234)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

func squareRoot(num int) (float64, error) {
	if num < 0 {
		return 0.0, fmt.Errorf("Input is a negative number. Cannot calculate square root")
	}

	return math.Sqrt(float64(num)), nil

}
