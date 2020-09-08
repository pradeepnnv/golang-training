package main

import "fmt"

func main() {
	myPerson := person{
		first: "First",
		last:  "Last",
		age:   4,
	}
	myPerson.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p)
}
