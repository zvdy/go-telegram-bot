package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	news, err := getDevNews()
	if err != nil {
		panic(err)
	}

	fmt.Println(news)
}

func getDevNews() (string, error) {
	// Scrape the developer news from DuckDuckGo on dev.to and return the top 5
	doc, err := goquery.NewDocument("https://duckduckgo.com/html/?q=developer+news")
	if err != nil {
		return "", err
	}

	var news []string
	doc.Find(".result__title").Each(func(i int, s *goquery.Selection) {
		if i < 7 {
			link := s.Find("a")
			title := link.Text()
			href, _ := link.Attr("href")
			news = append(news, fmt.Sprintf("%d. [%s](%s)", i+1, title, href))
		}
	}) // added closing parenthesis here

	if len(news) == 0 {
		return "No developer news found", nil
	}

	return fmt.Sprintf("Here are the top %d developer news:\n\n%s", len(news), strings.Join(news, "\n")), nil
}
