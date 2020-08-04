package main

import (
	"fmt"
	"os"
)

func main() {
	var answer1, answer2, answer3 string

	fmt.Print("Name:")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "scan-example: %s", err)
	}

	fmt.Print("Fav food:")
	_, err = fmt.Scan(&answer2)

	if err != nil {
		fmt.Fprintf(os.Stderr, "scan-example: %s", err)
	}

	fmt.Print("Fav sport:")
	_, err = fmt.Scan(&answer3)

	if err != nil {
		fmt.Fprintf(os.Stderr, "scan-example: %s", err)
	}

	fmt.Println(answer1, answer2, answer3)

}
