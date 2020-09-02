package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	v          vehicle
	fourWheels bool
}

type sedan struct {
	v      vehicle
	luxury bool
}

func main() {
	blueVehicle := vehicle{
		doors: 4,
		color: "blue",
	}

	greenVehicle := vehicle{
		doors: 2,
		color: "green",
	}
	myVan := truck{
		v:          blueVehicle,
		fourWheels: true,
	}
	mySedan := sedan{
		v:      greenVehicle,
		luxury: true,
	}
	fmt.Println(myVan)
	fmt.Println(mySedan)

	fmt.Println("Value of fourWheels for the object myTruck is :", myVan.fourWheels)
}
