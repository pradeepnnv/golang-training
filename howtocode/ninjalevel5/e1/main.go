package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		age   int
	}

	p := person{
		first: "Harry",
		last:  "Potter",
		age:   35,
	}

	fmt.Println(p)
}
