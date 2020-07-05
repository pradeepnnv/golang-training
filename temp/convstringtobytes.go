package main

import (
	"fmt"
)

func main() {
	s := `This is a string`
	bytesString := []byte(s)
	fmt.Printf("%T\t%T", s, bytesString)
	ns := string(bytesString)
	fmt.Printf("%T\t%T", bytesString, ns)
}
