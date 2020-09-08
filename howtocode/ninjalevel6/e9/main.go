package main

import "fmt"

func main() {
	f := func(s string) {
		fmt.Println("Value of s inside the func f is :", s)
	}

	foo(f)
}

func foo(f func(s string)) {
	f("Hello")
}
