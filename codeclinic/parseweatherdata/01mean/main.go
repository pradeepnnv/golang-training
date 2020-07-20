package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func main() {
	// f, err := os.Open("/Users/pnamburi/Downloads/Ex_Files_CodeClinic_Go/Exercise Files/01_parse_weather-data/Environmental_Data_Deep_Moor_2015.txt")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "readfiles: %s", err)
	// }

	resp, err := http.Get("https://raw.githubusercontent.com/lyndadotcom/LPO_weatherdata/master/Environmental_Data_Deep_Moor_2014.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "httpGet: %s", err)
	}

	rdr := csv.NewReader(resp.Body)
	rdr.TrimLeadingSpace = true
	rdr.Comma = '\t'
	rows, err := rdr.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "readfiles: %s", err)
	}

	fmt.Println("Number of records processed   :", len(rows)-1)
	fmt.Println("Mean Air Temperature          :", calculateMean(rows, 1))
	fmt.Println("Mean Barometric Pressure      :", calculateMean(rows, 2))
	fmt.Println("Mean Dew Point                :", calculateMean(rows, 4))

	fmt.Println("Median Air Temperature        :", calculateMedian(rows, 1))
	fmt.Println("Median Barometric Pressure    :", calculateMedian(rows, 2))
	fmt.Println("Median Dew Point              :", calculateMedian(rows, 4))
}

func calculateMean(rows [][]string, idx int) float64 {
	var sum float64
	var counter float64
	for i, row := range rows {
		if i != 0 {
			val, err := strconv.ParseFloat(row[idx], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "readfiles: Error converting airTemperature to float64: %s\n", err)
				continue
			}
			sum += val
			counter++
		}
	}
	return sum / counter
}

//Function to extract the sorted slice of all values at rowIdx and calculate median.
//If length of the slice is even, median is avg of
func calculateMedian(rows [][]string, rowIdx int) float64 {
	//create sorted slice of all values at idx
	var sorted []float64
	for i, row := range rows {
		if i > 0 {
			floatValue, err := strconv.ParseFloat(row[rowIdx], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "median: %s", err)
			}
			sorted = append(sorted, floatValue)
		}
	}
	sort.Float64s(sorted)

	//calculate median
	lenSorted := len(sorted)
	if lenSorted%2 == 0 {
		//length is even. So median is avg of middle numbers
		return (sorted[lenSorted/2] + sorted[lenSorted/2-1]) / 2
	}
	return sorted[lenSorted/2]
}
