package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

// main
func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		// fmt.Println(url)
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

// hitURL x
func hitURL(url string) error {
	resp, err := http.Get(url)
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Body)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(url, err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
