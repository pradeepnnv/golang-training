package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		"Micky",
		"Mouse",
		89,
	}
	p.speak()
}

func (p person) speak() {
	fmt.Printf("My name is %s,%s. My age is %d\n", p.first, p.last, p.age)
}
