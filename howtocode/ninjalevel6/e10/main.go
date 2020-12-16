package main

import "fmt"

func main() {
	var i int
	i = 25
	func() {
		fmt.Println(i)
	}()
}
