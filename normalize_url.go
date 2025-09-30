package main

import (
	net "net/url"
	"strings"
)

func normalizeURL(url string) (string, error) {
	urlObject, err := net.Parse(url)
	if err != nil {
		return "", err
	}
	normalizedUrl, err := net.JoinPath(urlObject.Hostname(), urlObject.Path)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(normalizedUrl, "/"), nil
}
