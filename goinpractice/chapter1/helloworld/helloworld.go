package main

import "fmt"

func greet(name string) string {
	return "Hey " + name + "!"
}
func main() {
	fmt.Println(greet("Pradeep"))
}
