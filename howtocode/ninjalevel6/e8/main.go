package main

import "fmt"

func main() {
	returnAFunc()()
}

func returnAFunc() func() {
	return func() {
		fmt.Println("This is an anonymous func being returned from returnAFunc()")
	}
}
