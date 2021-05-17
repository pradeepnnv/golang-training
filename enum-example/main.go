package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pradeepnnv/go-enum-example/car"
)

func main() {
	bmw := car.New(car.Honda, car.Red, "320i GT")
	ferrari := car.New(car.Ferrari, car.Red, "SF90 Stadala")

	fmt.Printf("I own a '%s' and dream about having a '%s'.\n", bmw, ferrari)

	cars := []*car.Car{bmw, ferrari}
	carsJSON, err := json.Marshal(cars)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(carsJSON))

	porscheJSON := []byte(`{"brand":"Porsche","color":"White","model":"Taycan"}`)
	var porsche car.Car

	err = json.Unmarshal(porscheJSON, &porsche)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Another car I love is ", porsche)
	cars = append(cars, &porsche)
	fmt.Println("Printing all my cars...", cars)

}
