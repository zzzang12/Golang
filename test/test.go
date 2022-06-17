package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	resp, err := http.Get("https://www.rettid.com/")
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(errRequestFailed)
	}
}
