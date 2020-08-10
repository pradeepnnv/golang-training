package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 7, 6, 4, 2, 3, 10, -1}
	go printCount(c)
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 1)
}
