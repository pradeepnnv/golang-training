package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	pMap := make(map[string]person)

	pMap["Harry"] = person{
		first: "Happy",
		last:  "Feet",
		age:   5,
	}
	fmt.Println(pMap)
}
