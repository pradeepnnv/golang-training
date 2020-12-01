package main

import "fmt"

func main() {
	mapFaves := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	mapFaves["hermione_granger"] = []string{`reading books`, `learning new things`, `dancing`}
	mapFaves["harry_potter"] = []string{`playing games`}

	for i, v := range mapFaves {
		fmt.Println(i, "\t", v)
	}
}
