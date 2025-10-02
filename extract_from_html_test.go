package main

import (
	"net/url"
	"reflect"
	"testing"
)

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
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getH1FromHTML(tc.html)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s\nExpected: %s\nActual: %s", i+1, tc.name, tc.expected, actual)
			}
		})
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
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getFirstParagraphFromHTML(tc.html)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s\nExpected: %s\nActual: %s", i+1, tc.name, tc.expected, actual)
			}
		})

	}
}

func TestGetURLsFromHTML(t *testing.T) {
	baseURL, _ := url.Parse("https://example.com")

	tests := []struct {
		name     string
		html     string
		baseURL  *url.URL
		expected []string
		wantErr  bool
	}{
		{
			name:    "Links with absolute URLs",
			html:    `<html><a href="https://example.com/page1">Link</a><img src="https://example.com/img1.jpg"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://example.com/page1",
			},
		},
		{
			name:    "Relative URLs resolved to absolute",
			html:    `<html><a href="/page2">Link</a><img src="images/photo.jpg"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://example.com/page2",
			},
		},
		{
			name:    "Mixed absolute and relative URLs",
			html:    `<html><a href="https://other.com/page">Link</a><img src="/img2.png"><a href="sub/page3">Link</a></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://other.com/page",
				"https://example.com/sub/page3",
			},
		},
		{
			name:     "Empty HTML",
			html:     "",
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:     "No URLs present",
			html:     `<html><p>Some text</p><div>No links here</div></html>`,
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:     "Malformed URLs",
			html:     `<html><a href="::invalid::">Link</a><img src="http://[invalid"></html>`,
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:     "Missing href attribute",
			html:     `<html><a>Link without href</a><img alt="no src"></html>`,
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:    "URLs with query parameters and fragments",
			html:    `<html><a href="/page4?q=1#section">Link</a><img src="https://example.com/img3.jpg?size=large"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://example.com/page4?q=1#section",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.html, tc.baseURL)
			if err != nil {
				t.Errorf("Test %v - %s\nUnexpected error: %v", i+1, tc.name, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s\nExpected: %v\nActual: %v", i+1, tc.name, tc.expected, actual)
			}
		})

	}
}

func TestGetImagesFromHTML(t *testing.T) {
	baseURL, _ := url.Parse("https://example.com")

	tests := []struct {
		name     string
		html     string
		baseURL  *url.URL
		expected []string
		wantErr  bool
	}{
		{
			name:    "Single image with absolute URL",
			html:    `<html><img src="https://example.com/img1.jpg"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://example.com/img1.jpg",
			},
			wantErr: false,
		},
		{
			name:    "Multiple images with relative URLs",
			html:    `<html><img src="/images/photo1.png"><img src="assets/img2.jpg"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://example.com/images/photo1.png",
				"https://example.com/assets/img2.jpg",
			},
		},
		{
			name:    "Mixed absolute and relative URLs",
			html:    `<html><img src="https://other.com/pic.jpg"><img src="/img3.png"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://other.com/pic.jpg",
				"https://example.com/img3.png",
			},
		},
		{
			name:     "Empty HTML",
			html:     "",
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:     "No images present",
			html:     `<html><p>Some text</p><a href="/link">Link</a></html>`,
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:     "Malformed src URL",
			html:     `<html><img src="::invalid::"></html>`,
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:     "Missing src attribute",
			html:     `<html><img alt="no src"></html>`,
			baseURL:  baseURL,
			expected: []string{},
		},
		{
			name:    "Image with query parameters",
			html:    `<html><img src="/img4.jpg?size=large"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://example.com/img4.jpg?size=large",
			},
		},
		{
			name:    "Image with attributes",
			html:    `<html><img src="/img5.gif" alt="test" class="image"></html>`,
			baseURL: baseURL,
			expected: []string{
				"https://example.com/img5.gif",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getImagesFromHTML(tc.html, tc.baseURL)
			if err != nil {
				t.Errorf("Test %v - %s\nUnexpected error: %v", i+1, tc.name, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s\nExpected: %v\nActual: %v", i+1, tc.name, tc.expected, actual)
			}
		})

	}
}

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		pageURL  string
		expected PageData
	}{
		{
			name:    "empty html",
			html:    "",
			pageURL: "http://example.com",
			expected: PageData{
				URL:            "http://example.com",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
		{
			name: "complete html with all elements",
			html: `
				<html>
					<head><title>Test</title></head>
					<body>
						<h1>Main Heading</h1>
						<p>First paragraph text.</p>
						<p>Second paragraph.</p>
						<a href="/about">About</a>
						<a href="https://external.com">External</a>
						<img src="image1.jpg">
						<img src="http://example.com/image2.png">
					</body>
				</html>`,
			pageURL: "http://example.com",
			expected: PageData{
				URL:            "http://example.com",
				H1:             "Main Heading",
				FirstParagraph: "First paragraph text.",
				OutgoingLinks:  []string{"http://example.com/about", "https://external.com"},
				ImageURLs:      []string{"http://example.com/image1.jpg", "http://example.com/image2.png"},
			},
		},
		{
			name: "missing elements",
			html: `
				<html>
					<body>
						<p>No heading here.</p>
						<a href="page.html">Link</a>
					</body>
				</html>`,
			pageURL: "http://example.com",
			expected: PageData{
				URL:            "http://example.com",
				H1:             "",
				FirstParagraph: "No heading here.",
				OutgoingLinks:  []string{"http://example.com/page.html"},
				ImageURLs:      []string{},
			},
		},
		{
			name: "relative and absolute URLs",
			html: `
				<html>
					<body>
						<h1>Heading</h1>
						<p>Text</p>
						<a href="/relative/path">Relative</a>
						<a href="subfolder/page.html">Subfolder</a>
						<a href="https://other.com">Absolute</a>
						<img src="/images/pic.jpg">
						<img src="local.jpg">
					</body>
				</html>`,
			pageURL: "http://example.com/path/",
			expected: PageData{
				URL:            "http://example.com/path/",
				H1:             "Heading",
				FirstParagraph: "Text",
				OutgoingLinks: []string{
					"http://example.com/relative/path",
					"http://example.com/path/subfolder/page.html",
					"https://other.com",
				},
				ImageURLs: []string{
					"http://example.com/images/pic.jpg",
					"http://example.com/path/local.jpg",
				},
			},
		},
		{
			name: "no body content",
			html: `
				<html>
					<head><title>Empty</title></head>
					<body></body>
				</html>`,
			pageURL: "http://example.com",
			expected: PageData{
				URL:            "http://example.com",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
		{
			name:    "invalid url",
			html:    "<html><h1>Test</h1><p>Paragraph</p></html>",
			pageURL: "not-a-url",
			expected: PageData{
				URL:            "not-a-url",
				H1:             "Test",
				FirstParagraph: "Paragraph",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.html, tc.pageURL)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Test %v - %s\nExpected: %+v\nActual: %+v", i+1, tc.name, tc.expected, actual)
			}
		})
	}
}
