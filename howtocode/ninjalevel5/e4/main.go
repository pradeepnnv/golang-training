package main

import "fmt"

func main() {
	s := struct {
		first string
		last  string
		age   int
	}{
		first: "Hua",
		last:  "Mulan",
		age:   16,
	}

	fmt.Println(s)
	fmt.Println(s.age)
}
