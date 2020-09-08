package main

import "fmt"

func main() {
	s := square{
		len: 2.98,
		wid: 4.96,
	}

	c := circle{
		r: 24.3233,
	}

	info(s)
	info(c)

}

type square struct {
	len float64
	wid float64
}

type circle struct {
	r float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (s square) area() float64 {
	return s.len * s.wid
}

func info(s shape) {
	fmt.Printf("Area of this shape is %f\n", s.area())
}
