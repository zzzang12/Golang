package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

var limit = 10
var baseURL = "https://kr.indeed.com/jobs?q=golang&limit=" + strconv.Itoa(limit) + "&start="

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + strconv.Itoa(limit*page)
	fmt.Println("Requesting", pageURL)
}

func getPages() (pages int) {
	res, err := http.Get(baseURL + "0")
	checkError(err)
	checkStatusCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %s", res.Status)
	}
}
