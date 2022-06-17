package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
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
		if err := hitURL(url); err != nil {
			fmt.Println(url)
			fmt.Println(err)
		}
	}
}

func hitURL(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
