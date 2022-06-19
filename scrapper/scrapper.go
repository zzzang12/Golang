package main

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Job struct {
	id              string
	title           string
	companyName     string
	companyLocation string
	snippet         string
}

var baseURL = "https://kr.indeed.com/jobs?q=golang&start="

func main() {
	var jobs []Job
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	printJobs(jobs)

	writeJobs(jobs)
}

func writeJobs(jobs []Job) {
	file, err := os.Create("jobs.csv")
	checkError(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Job Title", "Company Name", "Company Location", "Snippet"}
	err = w.Write(headers)
	checkError(err)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.companyName, job.companyLocation, job.snippet}
		err = w.Write(jobSlice)
		checkError(err)
	}
}

func getPage(page int) (jobs []Job) {
	pageURL := baseURL + strconv.Itoa(10*page)
	res, err := http.Get(pageURL)
	checkError(err)
	checkStatusCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".cardOutline").Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return
}

func printJobs(jobs []Job) {
	for _, job := range jobs {
		println(job.title, " / ", job.companyName, " / ", job.companyLocation)
		println(job.snippet)
	}
}

func extractJob(card *goquery.Selection) (job Job) {
	id, _ := card.Find(".jcs-JobTitle").Attr("data-jk")
	title := card.Find(".jcs-JobTitle>span").Text()
	companyName := card.Find(".companyName").Text()
	companyLocation := card.Find(".companyLocation").Text()
	snippet := strings.TrimSpace(card.Find(".job-snippet").Text())
	job = Job{id, title, companyName, companyLocation, snippet}
	return
}

func getPages() (pages int) {
	res, err := http.Get(baseURL)
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
