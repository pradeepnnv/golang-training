package main

import "fmt"

//Point struct represents a point
type Point struct {
	X int
	Y int
}

//Square struct
type Square struct {
	Center Point
	Length int
}

func (S *Square) move(dX int, dY int) {
	S.Center.X += dX
	S.Center.Y += dY
}

func main() {

	S1 := Square{
		Length: 5,
		Center: Point{
			X: 5,
			Y: 6,
		},
	}

	fmt.Println(S1)
	S1.move(5, 6)

	fmt.Println(S1)
	S1.move(1, 1)
	fmt.Println(S1)
}
