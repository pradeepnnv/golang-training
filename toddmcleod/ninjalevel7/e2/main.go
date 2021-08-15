package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
}

func main() {
	myPerson := person{
		"this is my name",
	}
	fmt.Println(myPerson)
	fmt.Println("Changing the value")
	changeMe(&myPerson)
	fmt.Println(myPerson)
}

func changeMe(p *person) {
	p.name = strings.ToUpper(p.name)
}
