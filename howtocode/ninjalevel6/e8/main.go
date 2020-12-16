package main

import "fmt"

func main() {
	x := returnAnonFunc()
	x()
}

func returnAnonFunc() func() {
	return func() {
		fmt.Println("This is an anonymous func")
	}
}
