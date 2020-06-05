//This program identifies the count of each words in the given String

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

	fmt.Println(text)
	words := strings.Fields(text)
	fmt.Println(words)
	wc := map[string]int{}
	fmt.Println(wc)
	for _, word := range words {
		wc[strings.ToLower(word)]++
	}
	fmt.Println(wc)

}
