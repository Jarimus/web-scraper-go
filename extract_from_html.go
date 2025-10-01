package main

import (
	"net/url"
	"strings"

	goquery "github.com/PuerkitoBio/goquery"
)

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
		return nil, err
	}
	// Find all links
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		urlString, exists := s.Attr("href")
		if !exists {
			return
		}
		newURL, err := url.Parse(urlString)
		if err != nil {
			return
		}
		if newURL.Hostname() == "" {
			newURL.Host = baseURL.Host
			newURL.Scheme = baseURL.Scheme
		}
		result = append(result, newURL.String())
	})

	return result, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	result := []string{}

	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	// Find all img srcs
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		urlString, exists := s.Attr("src")
		if !exists {
			return
		}
		newURL, err := url.Parse(urlString)
		if err != nil {
			return
		}
		if newURL.Hostname() == "" {
			newURL.Host = baseURL.Host
			newURL.Scheme = baseURL.Scheme
		}
		result = append(result, newURL.String())
	})

	return result, nil
}
