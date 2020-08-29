package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// DefaultReviewerCount contains the count of default reviewers for the repo
type DefaultReviewerCount struct {
	Size int `json:"size"`
}

func main() {
	start := time.Now()

	repoCount := 0

	repoFile, err := os.Open("repolist.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer repoFile.Close()

	baseURL := ""
	urlPath := "/default-reviewers"
	client := &http.Client{}

	if len(os.Args) < 2 {
		fmt.Println("AuthValue parameter is missing")
		os.Exit(1)
	}
	authValue := os.Args[1]

	repoInputScanner := bufio.NewScanner(repoFile)
	for repoInputScanner.Scan() {
		repoCount++
		repo := repoInputScanner.Text()
		// fmt.Println(repoInputScanner.Text())

		url := baseURL + repo + urlPath
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", "Basic "+authValue)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		// body, err := ioutil.ReadAll(resp.Body)

		// fmt.Println(string(body))
		reviewerCount := &DefaultReviewerCount{}
		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(reviewerCount); err != nil {
			fmt.Fprintf(os.Stderr, "bitbucketapiexample: %s", err)
		}
		// fmt.Println(reviewerCount.Size)
		fmt.Printf("Count of default reviewers for the repo %s is %d\n", repo, reviewerCount.Size)
	}
	fmt.Println("Time taken to identify default reviewers for ", repoCount, " is :", time.Now().Sub(start))

}