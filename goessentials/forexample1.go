package main

import (
	"fmt"
)

func main() {
	fmt.Println("Basic for example")
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("break  example")
	for i := 1; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("continue  example")
	for i := 1; i < 5; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}
}
