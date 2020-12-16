package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("This is main invoking foo")
}

func foo() {
	fmt.Println("This is foo")
}
