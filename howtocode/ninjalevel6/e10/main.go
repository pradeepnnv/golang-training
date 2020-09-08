package main

import "fmt"

func main() {
	i := 5
	foo(i)
	fmt.Println(i)
}

func foo(i int) {
	fmt.Println("Value of i in func foo is :", i)
	i = i + 5
	fmt.Println("Value of i in func foo after updating is ", i)
}
