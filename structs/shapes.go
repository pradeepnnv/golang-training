package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Length + r.Width)
// }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// func (c Circle) Perimeter() float64 {
// 	return 2 * math.Pi * c.Radius
// }

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
	// return 0
}
