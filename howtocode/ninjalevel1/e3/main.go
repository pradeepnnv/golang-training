package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x := 42
	y := "Aladdin"
	z := true
	str := fmt.Sprintf("%d, %s, %t", x, y, z)
	strv := fmt.Sprintf("%T, %T, %T", x, y, z)
	fmt.Println(str)
	fmt.Println(strv)
}
