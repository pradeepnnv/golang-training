package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "FirstName",
		last:  "Lastname",
		age:   24,
	}
	fmt.Println(p1)
}
