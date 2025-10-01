package main

import (
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
