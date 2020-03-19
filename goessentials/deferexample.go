package main

import (
	"fmt"
)

func main() {
	worker()
}

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")
	fmt.Println("worker")
}
