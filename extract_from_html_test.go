package main

import "testing"

func TestGetFirstHeaderFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "Single h1",
			html:     "<html><h1>Header One</h1></html>",
			expected: "Header One",
		},
		{
			name:     "Multiple headers",
			html:     "<html><h1>First Header</h1><h2>Second Header</h2></html>",
			expected: "First Header",
		},
		{
			name:     "No h1 present",
			html:     "<html><p>Some text</p><h2>Other Header</h2></html>",
			expected: "",
		},
		{
			name:     "Empty HTML",
			html:     "",
			expected: "",
		},
		{
			name:     "h1 with attributes",
			html:     "<html><h1 class=\"title\" id=\"main\">Main Header</h1></html>",
			expected: "Main Header",
		},
	}

	for i, tc := range tests {
		actual, err := getH1FromHTML(tc.html)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actual != tc.expected {
			t.Errorf("Test %v - %s\nExpected: %s\nActual: %s", i+1, tc.name, tc.expected, actual)
		}
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "One paragraph",
			html:     "<html><p>paragraph one</p></html>",
			expected: "paragraph one",
		},
		{
			name:     "Multiple paragraphs",
			html:     "<html><p>First paragraph</p><p>Second paragraph</p></html>",
			expected: "First paragraph",
		},
		{
			name:     "No paragraph present",
			html:     "<html><h1>Header</h1><div>Some text</div></html>",
			expected: "",
		},
		{
			name:     "Empty HTML",
			html:     "",
			expected: "",
		},
		{
			name:     "Paragraph with attributes",
			html:     "<html><p class=\"intro\" id=\"p1\">Intro paragraph</p></html>",
			expected: "Intro paragraph",
		},
	}

	for i, tc := range tests {
		actual, err := getFirstParagraphFromHTML(tc.html)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actual != tc.expected {
			t.Errorf("Test %v - %s\nExpected: %s\nActual: %s", i+1, tc.name, tc.expected, actual)
		}
	}
}
