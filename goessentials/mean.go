//Calculates mean of two floats
package main

import (
	"fmt"
)

func main() {

	x, y := 1.0, 29.45

	fmt.Printf("x=%f, type of %T\n", x, x)
	fmt.Printf("y=%f, type of %T\n", y, y)
	mean := (x + y) / 2.0
	fmt.Printf("result is %f, type of %T\n", mean, mean)
}
