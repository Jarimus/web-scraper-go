package main

import (
	"encoding/csv"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	writer.Comma = ';'

	// Write headers
	writer.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls", "references"})

	// For each page, write its data
	for _, data := range pages {
		record := []string{
			data.URL,
			data.H1,
			data.FirstParagraph,
			strings.Join(data.OutgoingLinks, ","),
			strings.Join(data.ImageURLs, ","),
		}
		err = writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
