package main

import "fmt"

func main() {
	anonFunc := func() {
		fmt.Println("This is an anonymous func!!!")
	}
	anonFunc()
}
