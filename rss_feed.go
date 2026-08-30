package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/kaszta1274/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (rssFeed *RSSFeed) unescapeFeed() {
	rssFeed.Channel.Title = unescapeHTML(rssFeed.Channel.Title)
	rssFeed.Channel.Description = unescapeHTML(rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = unescapeHTML(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = unescapeHTML(rssFeed.Channel.Item[i].Description)
	}
}

func unescapeHTML(text string) string {
	return html.UnescapeString(text)
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return nil, err
	}

	rssFeed.unescapeFeed()
	return &rssFeed, nil
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("couldn't get feed to fetch: %v", err)
		return
	}

	err = s.db.MarkFeedFetched(
		context.Background(),
		database.MarkFeedFetchedParams{
			ID:            feed.ID,
			LastFetchedAt: sql.NullTime{Time: time.Now().UTC()},
			UpdatedAt:     time.Now().UTC(),
		},
	)
	if err != nil {
		fmt.Printf("couldn't mark feed as fetched: %v", err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("couldn't fetch feed: %v", err)
		return
	}

	fmt.Printf("%s:\n", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	fmt.Println()
}
