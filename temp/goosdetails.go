//Program to print the OS details
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("GORoutines:\t", runtime.NumGoroutine())
	wg.Add(1)
	// time.Sleep(time.Second)
	go foo()
	fmt.Println("GORoutines:\t", runtime.NumGoroutine())
	// wg.Add()
	go bar()
	fmt.Println("GORoutines:\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("GORoutines:\t", runtime.NumGoroutine())
}

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
