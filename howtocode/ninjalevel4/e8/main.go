package main

import "fmt"

func main() {
	mymap := make(map[string][]string)
	mymap["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	mymap["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	mymap["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for k, v := range mymap {
		fmt.Println("Key is :", k)
		for _, s := range v {
			fmt.Println("\tValues are ", s)
		}
	}
}
