package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	go func() {
		defer wg.Done()
		fmt.Println("This is me printing from a goroutine")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("This is me printing from a goroutine again!!!")
	}()

	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
}
