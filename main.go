package main

import (
	"RestAPIClientInterface/RestAPIClient"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	url, _ := url.Parse("https://www.google.com")
	req := &http.Request{
		Method: "GET",
		URL:    url,
	}
	resp, err := RestAPIClient.WebCallV1(req)
	if err == nil && resp != nil {
		fmt.Println(resp.StatusCode)
	} else {
		fmt.Println(err)
	}
}
