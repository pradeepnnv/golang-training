package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go printSomething()
	go printSomething()
	wg.Wait()
}

func printSomething() {
	fmt.Println("Printing something")
	wg.Done()
}
