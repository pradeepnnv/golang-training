package main

import "fmt"

// func main() {
// 	fmt.Println(Hello())
// }

const englishHelloPrefix = "Hello "

func Hello(name string) string {
	return fmt.Sprintf("%s%s!!", englishHelloPrefix, name)
}
