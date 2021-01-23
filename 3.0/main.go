package main

import (
	"errors"
	"fmt"
	"net/http"
)

// errHttpNotFound
var errHttpNotFound = errors.New("http not found.")

// main
func main() {
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
		"https://mgkim.cafe24.com/",
		"http://www.xxdee.dfd.com/",
	}
	for _, url := range urls {
		fmt.Println(url)
		hitUrl(url)
	}
}

// hitUrl
func hitUrl(url string) error {
	resp, err := http.Get(url)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	if err != nil || resp.StatusCode >= 400 {
		return errHttpNotFound
	}
	return nil
}
