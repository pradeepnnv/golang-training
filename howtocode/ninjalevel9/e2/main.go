package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("my name is :" + p.name)
}

type human interface {
	speak()
}

func (p person) saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		"thing1",
	}
	// p.speak()
	p.saySomething(&p)
}
