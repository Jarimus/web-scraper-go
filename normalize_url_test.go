package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme (https)",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme (http)",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove final forward slash (https)",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove final forward slah (http)",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove port, queries, sections etc",
			inputURL: "http://blog.boot.dev:1234/path/?query=1&search=stuff#section",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("\nTest %v - '%s' \nunexpected error: %v", i+1, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("\nTest %v - %s \ninput: %v\nexpected: %v,\nactual: %v", i+1, tc.name, tc.inputURL, tc.expected, actual)
			}
		})
	}
}
