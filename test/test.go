package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	url    string
	status string
}

func main() {
	channel := make(chan Result)

	go hitURL("https://jerryjerryjerry.tistory.com/99", channel) // resp ok err nil
	go hitURL("https://www.rettid.com/", channel)

	for i := 0; i < 2; i++ {
		c := <-channel
		fmt.Println(c.url, c.status)
	}
}

func hitURL(url string, channel chan<- Result) {
	fmt.Println("Checking:", url)
	status := ""
	if resp, err := http.Get(url); resp == nil {
		status = "wrong address"
	} else if err != nil {
		status = "false"
	} else {
		status = "ok"
	}
	channel <- Result{url, status}
}
