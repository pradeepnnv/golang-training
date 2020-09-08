package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello. This is an anonymous func")
	}()
}
