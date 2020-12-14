package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "test.txt"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error while opening the file: \n", err)
	}
	inputScanner := bufio.NewScanner(f)

	wMap := make(map[string]int)

	for inputScanner.Scan() {
		line := inputScanner.Text()
		words := strings.Split(line, " ")
		for _, w := range words {
			if w != "" {
				wMap[w]++
			}
		}
	}

	fmt.Println(wMap)

}
