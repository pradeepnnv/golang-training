package main

import "fmt"

func main() {

	var h human
	p := person{
		age:  25,
		name: "Moriarty Holmes",
	}
	h = p
	fmt.Println(h.speak())
	saySomething(h)

}

type person struct {
	age  int
	name string
}

func (p person) speak() string {
	return fmt.Sprintf("My name is %s and I'm %d years old. Spiffing!!!", p.name, p.age)
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}
