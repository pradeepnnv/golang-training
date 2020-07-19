package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.Open("/Users/pnamburi/Downloads/Ex_Files_CodeClinic_Go/Exercise Files/01_parse_weather-data/Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "median: %s", err)
	}
	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rows, err := rdr.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "median: %s", err)
	}

	fmt.Println("Number of rows processed:", len(rows)-1)
	// if rowLength%2 != 0 {
	// 	median, err = strconv.ParseFloat(rows[rowLength/2+1][2], 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "median: %s", err)
	// 	}
	// } else {
	// 	median1, err := strconv.ParseFloat(rows[rowLength/2-1][2], 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "median: %s", err)
	// 	}
	// 	median2, err := strconv.ParseFloat(rows[rowLength/2+1][2], 64)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "median: %s", err)
	// 	}
	// 	median = (median1 + median2) / 2
	// }
	medianAirTemp := calculateMedian(rows, 1)
	fmt.Println("Median Air Temperature:", medianAirTemp)

	fmt.Println("Elapsed time:", time.Now().Sub(start))
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
