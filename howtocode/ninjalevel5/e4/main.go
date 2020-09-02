package main

import "fmt"

func main() {
	anonymousStruct := struct {
		name string
	}{
		"pradeep",
	}
	fmt.Println(anonymousStruct)
}
