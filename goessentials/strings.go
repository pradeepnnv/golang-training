//Explores various string functions

package main

import (
	"fmt"
)

func main() {
	book := "The color of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	fmt.Println(book[4:11])
	fmt.Println(book[:2])
	fmt.Println(book[6:])

	fmt.Println("t" + book[1:])

	fmt.Println("it was 1/2 price")
	poem := `
                The road goes ever on
                Down from the door where it began
                ...
                `
	fmt.Println(poem)
}
