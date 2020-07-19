package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("/Users/pnamburi/Downloads/Ex_Files_CodeClinic_Go/Exercise Files/01_parse_weather-data/Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "readfiles: %s", err)
	}
	// inputScanner := bufio.NewScanner(f)
	// for inputScanner.Scan() {
	// 	fmt.Println(inputScanner.Text())
	// }
	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rows, err := rdr.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "readfiles: %s", err)
	}

	var airTemp, counter float64 //Mean Air Temperature

	for i, row := range rows {
		if i != 0 {
			aT, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "readfiles: Error converting airTemperature to float64: %s\n", err)
				continue
			}
			airTemp += aT
			counter++
		}
	}
	fmt.Println("Total number of records:", counter)
	fmt.Println("Mean Air Temperature: ", airTemp/counter)
}
