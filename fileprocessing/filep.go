package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("/tmp/hello-world-service-request.log")
	defer inputFile.Close()
	if err != nil {
		fmt.Println("ERR:", err)
	}
	inputScanner := bufio.NewScanner(inputFile)

	outputFile, err := os.Create("logdetails.txt")
	defer outputFile.Close()
	if err != nil {
		fmt.Println("ERR:", err)
	}

	for inputScanner.Scan() {
		// fmt.Println(inputScanner.Text())
		logLine := inputScanner.Text()
		logEntries := strings.Split(logLine, " ")
		fmt.Println(logEntries[5])
		io.WriteString(outputFile, fmt.Sprintf("%s %s %s\n", logEntries[5], logEntries[6], logEntries[7]))
	}
}
