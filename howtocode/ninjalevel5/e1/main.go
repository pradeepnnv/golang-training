package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	iceCreamFlavours []string
}

func main() {
	p1 := person{
		firstName:        "Bees",
		lastName:         "Wax",
		iceCreamFlavours: []string{"Vanilla", "Strawberry"},
	}
	p2 := person{
		firstName:        "Austin",
		lastName:         "Powers",
		iceCreamFlavours: []string{"Redvelet", "Currant"},
	}
	fmt.Println(p1, p2)
}
