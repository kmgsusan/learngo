package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
	}

	c := make(chan string)
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

// hitURL x
func hitURL(url string, c chan string) {
	resp, err := http.Get(url)
	result := "OK"
	if err != nil || resp.StatusCode >= 400 {
		result = "FAILD"
	}
	c <- url + " : " + result + "(" + strconv.Itoa(resp.StatusCode) + ")"
}
