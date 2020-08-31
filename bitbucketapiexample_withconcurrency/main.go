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
	ch := make(chan string)
	repoInputScanner := bufio.NewScanner(repoFile)
	for repoInputScanner.Scan() {
		repoCount++
		repo := repoInputScanner.Text()
		// fmt.Println(repoInputScanner.Text())
		go getDefaultReviewerCount(repo, ch)
	}
	for i := 0; i < repoCount; i++ {
		fmt.Println(<-ch) //receive from channel
	}
	fmt.Println("Time taken to identify default reviewers for ", repoCount, " is :", time.Now().Sub(start))

}

func getDefaultReviewerCount(repo string, ch chan<- string) {

	urlPath := "/default-reviewers"
	client := &http.Client{}

	if len(os.Args) < 3 {
		fmt.Println("AuthValue or BaseURL parameters are missing")
		fmt.Println("Usage: run main.go AuthValue BaseURL")
		os.Exit(1)
	}
	authValue := os.Args[1]
	baseURL := os.Args[2]
	url := baseURL + repo + urlPath

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Basic "+authValue)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- fmt.Sprintf("Got an invalid response code of %d while processing the repo %s. Please verify.\n", resp.StatusCode, repo)
		return
	}

	reviewerCount := &DefaultReviewerCount{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(reviewerCount); err != nil {
		fmt.Fprintf(os.Stderr, "bitbucketapiexample: %s", err)
	}
	ch <- fmt.Sprintf("Count of default reviewers for the repo %s is %d\n", repo, reviewerCount.Size)

}
