package main

import (
	"fmt"
	"math"
)

func squareroot(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("Error while calculating square root of a negative number (%f)", n)
	}
	return math.Sqrt(n), nil
}

func main() {
	r, err := squareroot(4.0)
	if err != nil {
		fmt.Printf("ERROR:%s", err)
	} else {
		fmt.Println(r)
	}
	r, err = squareroot(-3)
	if err != nil {
		fmt.Printf("Error:%s", err)
	} else {
		fmt.Println(r)
	}
}
