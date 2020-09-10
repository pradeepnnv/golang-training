package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go printSomething(&wg)
		go printSomething(&wg)
	}
	wg.Wait()
}

func printSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Something")
}
