package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		fmt.Println("Usage: <url> <max concurrency> <max pages to crawl>")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		fmt.Println("Usage: <url> <max concurrency> <max pages to crawl>")
		os.Exit(1)
	}
	rawBaseURL := os.Args[1]
	maxConcurrency := 3
	maxPages := 25
	if len(os.Args) == 3 {
		var err error
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Please provide a valid number for concurrency.")
			os.Exit(1)
		}
	} else if len(os.Args) == 4 {
		var err error
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Please provide a valid number for concurrency.")
			os.Exit(1)
		}
		maxPages, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Please provide a valid number for concurrency.")
			os.Exit(1)
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

	for normalizedURL, pageData := range cfg.pages {
		fmt.Printf("%d - %s\n", pageData.Visits, normalizedURL)
	}
	fmt.Printf("Pages crawled: %d\n", len(cfg.pages))
}
