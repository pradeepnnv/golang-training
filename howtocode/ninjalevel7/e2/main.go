package main

import "fmt"

func main() {
	p1 := person{
		"Boo",
		"Yeah",
		29,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func changeMe(newp *person) {
	newp.age = 24
}
