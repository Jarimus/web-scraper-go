package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const filenameCSV = "report.csv"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no website provided\nUsage: <url> <max concurrency> <max pages to crawl>")
	}
	if len(os.Args) > 4 {
		log.Fatal("too many arguments provided\nUsage: <url> <max concurrency> <max pages to crawl>")
	}
	rawBaseURL := os.Args[1]
	maxConcurrency := 3
	maxPages := 1000
	if len(os.Args) == 3 {
		var err error
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Please provide a valid number for concurrency.")
		}
	} else if len(os.Args) == 4 {
		var err error
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("Please provide a valid number for concurrency.")
		}
		maxPages, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal("Please provide a valid number for concurrency.")
		}
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\nConcurrency: %d\nMax pages: %d\n", rawBaseURL, maxConcurrency, maxPages)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	err = writeCSVReport(cfg.pages, filenameCSV)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for normalizedURL, pageData := range cfg.pages {
		fmt.Printf("%d - %s\n", pageData.Visits, normalizedURL)
	}
	fmt.Printf("Pages crawled: %d\n", len(cfg.pages))
}
