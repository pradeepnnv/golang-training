package main

import "fmt"

func main() {
	var i int
	i = 1
	for {
		fmt.Println(i)
		if i >= 35 {
			break
		}
		i++
	}
}
