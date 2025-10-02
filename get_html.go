package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", nil
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// Check status code
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("error (%d) getting %s", res.StatusCode, rawURL)
	}
	// Check Content-Type (text/html)
	if !strings.Contains(res.Header.Get("Content-Type"), "text/html") {
		return "", fmt.Errorf("content-type is not text/html fo %s", rawURL)
	}
	return string(body), nil
}
