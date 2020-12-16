package main

import "fmt"

func main() {
	c := circle{
		24.8,
	}
	s := square{
		19.6,
		29.4,
	}
	info(s)
	info(c)
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return (c.radius * c.radius) * (22 / 7)
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("Area of the shape is %f\n", s.area())
}
