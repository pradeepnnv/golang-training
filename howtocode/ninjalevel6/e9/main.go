package main

import "fmt"

func main() {
	testFunc(func() {
		fmt.Println("Hi")
	})
}

func testFunc(x func()) {
	x()
}
