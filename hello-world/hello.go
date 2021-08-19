package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const suffix = "!!"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return fmt.Sprintf("%s%s%s", spanishHelloPrefix, name, suffix)
	}
	if language == french {
		return fmt.Sprintf("%s%s%s", frenchHelloPrefix, name, suffix)
	}
	return fmt.Sprintf("%s%s%s", englishHelloPrefix, name, suffix)
}
