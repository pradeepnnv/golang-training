package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("This is an anonymous func assigned to a variable")
	}
	f()
}
