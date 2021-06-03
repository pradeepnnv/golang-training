package main

import "fmt"

func main() {
	name := getName()
	fmt.Println("Name is :", name)
}

func getName() string {
	return "Hello World!"
}
