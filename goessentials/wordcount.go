//wordcount calculates the count of words in a string
package main

import (
	"fmt"
	"strings"
)

func main() {
	wordcount := map[string]int{}
	fmt.Println(wordcount)

	text := `
Needles and pins
Needles and pins
Sew me a sail
To Catch me the wind
  `
	// fmt.Println(text)
	//
	// fmt.Println(strings.Fields(text))

	wordlist := strings.Fields(text)
	for _, eachword := range wordlist {
		// fmt.Println(strings.ToLower(eachword))
		// eachword = strings.ToLower(eachword)
		// count := 0
		// _, ok := wordcount[eachword]
		// if ok {
		// 	count = wordcount[eachword]
		// 	count++
		// 	wordcount[eachword] = count
		// } else {
		// 	wordcount[eachword] = 1
		// }
		wordcount[strings.ToLower(eachword)]++
	}
	fmt.Println(wordcount)
}
