package main

import (
	"encoding/json"
	"fmt"
)

type Battery struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Description Info   `json:"info"`
}

type Info struct {
	Description string `json:"desc"`
}

func main() {
	jsonString := `{"name":"battery console","capacity":40,"time":"June 30, 2022 at 5:17 PM",
	"info":{"desc":"this is battery reading"}
	}`
	fmt.Println("json string is :", jsonString)
	var battery Battery

	err := json.Unmarshal([]byte(jsonString), &battery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(battery.Description)

}
