package main

import (
	"fmt"
	net "net/url"
	"strings"
)

func normalizeURL(url string) (string, error) {
	urlObject, err := net.Parse(url)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", err)
	}
	normalizedUrl, err := net.JoinPath(urlObject.Hostname(), urlObject.Path)
	if err != nil {
		return "", fmt.Errorf("error joining host and path: %v", err)
	}
	normalizedUrl = strings.ToLower(normalizedUrl)
	return strings.TrimRight(normalizedUrl, "/"), nil
}
