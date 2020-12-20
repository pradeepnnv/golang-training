package main

import "fmt"

type person struct {
	age   int
	first string
	last  string
}

func main() {
	p := person{
		15,
		"Donald",
		"Duck",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	// (*p).age = 24
	// (*p).first = "First"
	// (*p).last = "Last"
	p.age = 24
	p.first = "First"
	p.last = "Last"
}
