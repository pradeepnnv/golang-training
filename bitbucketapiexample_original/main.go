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

	urlPath := "/default-reviewers"
	client := &http.Client{}

	if len(os.Args) < 3 {
		fmt.Println("AuthValue or BaseURL parameters are missing")
		fmt.Println("Usage: run main.go AuthValue BaseURL")
		os.Exit(1)
	}
	authValue := os.Args[1]
	baseURL := os.Args[2]

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
		// fmt.Println(resp.StatusCode)
		// body, err := ioutil.ReadAll(resp.Body)

		// fmt.Println(string(body))
		if resp.StatusCode != 200 {
			fmt.Fprintf(os.Stderr, "Got an invalid response code of %d while processing the repo %s. Please verify.\n", resp.StatusCode, repo)
			continue
		}
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
