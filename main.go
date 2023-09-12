package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(botToken, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if strings.HasPrefix(update.Message.Text, "/news") {
		news, err := getDeveloperNews()
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Error getting developer news",
			})
			return
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   news,
		})
	}
}

func getDeveloperNews() (string, error) {
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
