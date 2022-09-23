package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// get URL from command line parameters
	url := os.Args[1]

	// if no parameter warn user
	if url == "" {
		log.Fatal("Please provide a URL")
	}

	f, err := os.OpenFile("results.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// read the response body
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	log.Println("response Status:", res.Status)

	// check the content of the response if it includes any links
	// if it does, then create a slice of all links
	// if it doesn't, then return an empty slice
	links := getLinks(res.Body)

	// print the links
	for _, link := range links {
		log.Println("--------------------")
		log.Print(link)

		// visit each URL and report the status code
		req, err := http.NewRequest("GET", link, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		// print the status code
		log.Println(res.Status)
		defer res.Body.Close()
	}

}

// getLinks reads the content of the response body and returns a slice of all links
func getLinks(body io.Reader) []string {
	links := []string{}

	// create a new goquery document from the HTTP response
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		link, _ := s.Attr("href")
		// iff link starts with http, then append it to the slice
		if len(link) > 4 && link[:4] == "http" {
			links = append(links, link)
		}
	})

	return links
}
