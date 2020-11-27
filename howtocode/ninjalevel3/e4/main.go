package main

import "fmt"

func main() {
	i := 1
	for {
		if i < 34 {
			i++
			fmt.Println(i)
		} else {
			break
		}
	}
}
