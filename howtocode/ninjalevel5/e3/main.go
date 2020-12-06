package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	v      vehicle
	luxury bool
}

type truck struct {
	v          vehicle
	fourWheels bool
}

func main() {

	dodgeRAM := truck{
		v: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheels: true,
	}

	chevyCoupe := sedan{
		v: vehicle{
			doors: 2,
			color: "periwinkle",
		},
		luxury: true,
	}

	fmt.Println(dodgeRAM)
	fmt.Println(chevyCoupe)
}
