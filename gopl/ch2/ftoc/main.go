//Ftoc converts temperature from Fahrenheit to Celsius
package main

import "fmt"

func main() {
	const freeezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freeezingF, ftoc(freeezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, ftoc(boilingF))
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
