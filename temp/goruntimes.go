package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUS: ", runtime.NumCPU())
	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()

		}()
		fmt.Println("Number of go routines:", runtime.NumGoroutine())
	}
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)

}
