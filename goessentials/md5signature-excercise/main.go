package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// h := md5.New()
	// io.WriteString(h, "This is a test")
	// // fmt.Println(h.Sum(nil))
	// fmt.Printf("%x \n", h.Sum(nil))

	// dat, err := ioutil.ReadFile("./main.go")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "md5signature-excersice: %s", err)
	// }
	// fmt.Printf("md5sum of ./main.go is %x\n", md5.Sum(dat))

	f, err := os.Open("/Users/pnamburi/Downloads/Ex_Files_Go_EssT/Exercise Files/Ch06/06_07/nasa-logs/md5sum.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "md5signature-excercise: %s", err)
		// log.Fatalf("failed opening the file: %s", err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, " ")
		// fmt.Printf("MD5 Hash of file %s is %s\n", val[2], val[0])
		dat, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", "/Users/pnamburi/Downloads/Ex_Files_Go_EssT/Exercise Files/Ch06/06_07/nasa-logs", val[2]))
		if err != nil {
			// log.Fatalf("failed opening the file: %s", err)
			fmt.Fprintf(os.Stderr, "md5signature-excercise: %s", err)
		}
		md5Hash := fmt.Sprintf("%x", md5.Sum(dat))
		// fmt.Println(md5Hash)
		if md5Hash != val[0] {
			fmt.Printf("Md5 hash of file %s is not matching.\n", val[2])
		} else {
			fmt.Printf("All is well for %s\n", val[2])
		}

	}

}
