//Program to identify the count of times a word appears in a text

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
  Needles and pins
  Needles and pins
  Sew me a sail
  To catch me the wind`
	// fmt.Println(text)

	words := strings.Fields(text)
	wc := make(map[string]int)
	// fmt.Println(words)
	for _, s := range words {
		fmt.Println(strings.ToLower(s))
		wc[strings.ToLower(s)]++
	}
	fmt.Println(wc)
}
