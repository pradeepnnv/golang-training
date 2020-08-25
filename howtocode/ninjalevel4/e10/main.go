package main

import "fmt"

func main() {
	mymap := make(map[string][]string)
	mymap["bond_james"] = []string{`Shaken, not stirred`, `Martinis`, `Women`}
	mymap["moneypenny_miss"] = []string{`James Bond`, `Literature`, `Computer Science`}
	mymap["no_dr"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	fmt.Println(mymap)
	mymap["pradeep"] = []string{"Reading books", "playing with kids", "working out"}
	fmt.Println(mymap)

	delete(mymap, "pradeep")
	fmt.Println(mymap)
}
