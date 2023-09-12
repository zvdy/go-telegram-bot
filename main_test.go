package main

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

// Mock the goquery.NewDocument function
var getDocument = goquery.NewDocument

func TestGetDeveloperNews(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name: "returns developer news",
			expected: `Here are the top 7 developer news:

1. [The 2021 State of Developer Ecosystem by JetBrains](https://www.jetbrains.com/lp/devecosystem-2021/)
2. [The 2021 State of Developer Ecosystem by JetBrains](https://www.jetbrains.com/lp/devecosystem-2021/)
3. [The 2021 State of Developer Ecosystem by JetBrains](https://www.jetbrains.com/lp/devecosystem-2021/)
4. [The 2021 State of Developer Ecosystem by JetBrains](https://www.jetbrains.com/lp/devecosystem-2021/)
5. [The 2021 State of Developer Ecosystem by JetBrains](https://www.jetbrains.com/lp/devecosystem-2021/)
6. [The 2021 State of Developer Ecosystem by JetBrains](https://www.jetbrains.com/lp/devecosystem-2021/)
7. [The 2021 State of Developer Ecosystem by JetBrains](https://www.jetbrains.com/lp/devecosystem-2021/)`,
		},
		{
			name:     "returns error when unable to scrape developer news",
			expected: "Error getting developer news",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "returns error when unable to scrape developer news" {
				// Mock the goquery.NewDocument function to return an error
				getDocument = func(url string) (*goquery.Document, error) {
					return nil, assert.AnError
				}
				defer func() { getDocument = goquery.NewDocument }()
			} else {
				// Mock the goquery.NewDocument function to return a document with developer news
				getDocument = func(url string) (*goquery.Document, error) {
					html := `<html><body><div class="result__title"><a href="https://www.jetbrains.com/lp/devecosystem-2021/">The 2021 State of Developer Ecosystem by JetBrains</a></div></body></html>`
					return goquery.NewDocumentFromReader(strings.NewReader(html))
				}
				defer func() { getDocument = goquery.NewDocument }()
			}

			actual, err := getDeveloperNews()
			if tt.expected == "Error getting developer news" {
				assert.Error(t, err)
				assert.Equal(t, tt.expected, actual)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
