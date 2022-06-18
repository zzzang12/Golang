package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

type ExtractedJob struct {
	id              string
	jobTitle        string
	companyName     string
	companyLocation string
	snippet         string
}

var baseURL = "https://kr.indeed.com/jobs?q=golang&start="

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + strconv.Itoa(10*page)
	fmt.Println("Requesting", pageURL)

	res, err := http.Get(pageURL)
	checkError(err)
	checkStatusCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".cardOutline").Each(func(i int, card *goquery.Selection) {
		id, _ := card.Find(".jcs-JobTitle").Attr("data-jk")
		jobTitle := card.Find(".jcs-JobTitle>span").Text()
		companyName := card.Find(".companyName").Text()
		companyLocation := card.Find(".companyLocation").Text()
		snippet := card.Find(".job-snippet").Text()
		extractedJob := ExtractedJob{id, jobTitle, companyName, companyLocation, snippet}
		fmt.Println(extractedJob.jobTitle, "/", extractedJob.companyName, "/", extractedJob.companyLocation)
		fmt.Println(extractedJob.snippet)
	})
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
