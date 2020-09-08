package main

import "fmt"

func main() {
	defer bar()
	fmt.Println("Inside main")
}

func bar() {
	fmt.Println("Inside bar")
}
