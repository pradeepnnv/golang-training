package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(boo())
}

func foo() int {
	return 5
}

func boo() (int, string) {
	return 0, "error"
}
