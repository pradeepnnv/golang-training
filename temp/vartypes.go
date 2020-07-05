package main

import (
	"fmt"
)

func main() {
	var boolVar bool
	var intVar int
	var strVar string

	boolVar = true
	intVar = 25
	strVar = `This is a "quoted" string`

	fmt.Printf("Bool var is of type %T and value is ====>%t\n", boolVar, boolVar)
	fmt.Printf("Int var is of type %T and value is =====>%d\n", intVar, intVar)
	fmt.Printf("String var is of type %T and value is ==>%s\n", strVar, strVar)

	//Converting string into a slice of bytes
	strBytes := ([]byte)(strVar)

	fmt.Printf("String converted into a byte slice is of type %T and value is %v\n", strBytes, strBytes)

	//Converting string into a slice of runes
	strRunes := ([]rune)(strVar)

	fmt.Printf("String converted into a rune slice is of type %T and value is %v\n", strRunes, strRunes)
}
