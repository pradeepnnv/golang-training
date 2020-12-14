package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(boo())
}

func foo() int {
	return 1
}

func boo() (int, string) {
	return 1, "one"
}
