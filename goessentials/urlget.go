//urlget prints the http header Content-Type in the response while performing GET on a URL

package main

import (
	"fmt"
	"net/http"
)

func main() {
	header, err := contentType("https://api.kube.hsp-dev.telematics.cccis.com")
	if err != nil {
		fmt.Printf("ERROR:%s", err)
	} else {
		fmt.Println(header)
	}
}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error while retreiving Content Type header when GET is performed on %s is (%f)", url, err)
	}
	return resp.Header.Get("Content-Type"), nil
}
