//Program to print the OS details
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	foo()
	bar()
	fmt.Println("OS:\t", runtime.GOOS)
	fmt.Println("ARCH:\t", runtime.GOARCH)
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GORoutines:\t", runtime.NumGoroutine())
	time.Sleep(time.Second)
}

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
