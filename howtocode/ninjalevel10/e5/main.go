package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1
	}()
	v, ok := <-c

	fmt.Println("Before closing the channel:", v, ok)

	close(c)

	v, ok = <-c
	fmt.Println("After closing the channel", v, ok)
}
