package main

import (
	"fmt"
	"net/url"
	"strings"

	goquery "github.com/PuerkitoBio/goquery"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
	Visits         int
}

func extractPageData(html, pageURL string) PageData {
	// Get <h1>
	h1, err := getH1FromHTML(html)
	if err != nil {
		fmt.Printf("Error getting <h1>: %s", err.Error())
	}
	// Get first paragraph
	p1, err := getFirstParagraphFromHTML(html)
	if err != nil {
		fmt.Printf("Error getting first paragraph: %s", err.Error())
	}
	// Parse url string into a url object
	baseUrl, err := url.Parse(pageURL)
	if err != nil {
		fmt.Printf("Error parsing pageURL: %s", err.Error())
		return PageData{
			URL:            pageURL,
			H1:             h1,
			FirstParagraph: p1,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
			Visits:         1,
		}
	}
	// Get outgoing links
	outgoingLinks, err := getURLsFromHTML(html, baseUrl)
	if err != nil {
		fmt.Printf("Error getting outgoing links: %s", err.Error())
	}
	// Get img urls
	imgURLS, err := getImagesFromHTML(html, baseUrl)
	if err != nil {
		fmt.Printf("Error getting images: %s", err.Error())
	}
	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: p1,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imgURLS,
		Visits:         1,
	}
}

func getH1FromHTML(html string) (string, error) {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}
	return doc.Find("h1").First().Text(), nil
}

func getFirstParagraphFromHTML(html string) (string, error) {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}
	return doc.Find("p").First().Text(), nil
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	result := []string{}

	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []string{}, err
	}
	// Find all links
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		urlString, exists := s.Attr("href")
		if !exists {
			return
		}
		newURL, err := url.Parse(urlString)
		if err != nil {
			fmt.Printf("couldn't parse src %q: %v\n", urlString, err)
			return
		}
		absoluteURL := baseURL.ResolveReference(newURL)
		result = append(result, absoluteURL.String())
	})

	return result, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []string{}, err
	}

	// Find all img srcs
	result := []string{}
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		urlString, exists := s.Attr("src")
		if !exists || strings.TrimSpace(urlString) == "" {
			return
		}
		newURL, err := url.Parse(urlString)
		if err != nil {
			fmt.Printf("couldn't parse src %q: %v\n", urlString, err)
			return
		}
		absoluteURL := baseURL.ResolveReference(newURL)
		result = append(result, absoluteURL.String())
	})

	return result, nil
}
