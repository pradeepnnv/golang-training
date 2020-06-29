//Program to process request logs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("/tmp/hello-world-service-request.log")
	if err != nil {
		fmt.Println("ERR:", err)
	}
	inputFileScanner := bufio.NewScanner(inputFile)

	for inputFileScanner.Scan() {
		fmt.Println(inputFileScanner.Text())
	}
}
