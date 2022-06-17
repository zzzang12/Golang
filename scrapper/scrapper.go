package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	channel := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.rettid.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for range urls {
		c := <-channel
		fmt.Println(c.url, c.status)
	}
}

func hitURL(url string, channel chan<- requestResult) {
	status := ""
	if resp, err := http.Get(url); resp == nil {
		status = "wrong address"
	} else if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	channel <- requestResult{url, status}
}
