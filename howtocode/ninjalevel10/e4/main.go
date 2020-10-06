package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	//convert into anonmyous func
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

//func to receive from multiple channels through case statements
func receive(c <-chan int, q <-chan int) {
	//infinite for loop
	for {
		//select without a statement
		select {
		case v := <-c:
			{
				fmt.Println(v)
			}
		case <-q:
			return
		}
	}

}
