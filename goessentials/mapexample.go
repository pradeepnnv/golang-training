//This program demonstrates usage of maps in Go.

package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 11234.23,
		"GOOG": 113232.234,
	}
	fmt.Println(stocks)
}
