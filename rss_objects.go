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

	"github.com/theblakeyg/blog-aggregator/internal/database"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	//build request
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	request.Header.Add("User-Agent", "gator")

	//build client
	client := http.Client{Timeout: http.DefaultClient.Timeout}

	//send the request
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//turn the response body into a byte array
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	result := &RSSFeed{}

	//turn the XML into our struct
	err = xml.Unmarshal(data, result)
	if err != nil {
		return &RSSFeed{}, err
	}

	result.Channel.Title = html.UnescapeString(result.Channel.Title)
	result.Channel.Description = html.UnescapeString(result.Channel.Description)

	for _, item := range result.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return result, nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.database.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed: %v", err)
	}

	args := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            feed.ID,
	}

	err = s.database.MarkFeedFetched(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error marking feed as feteched: %v", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %v", err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("|")
	fmt.Printf("| %v\n", rssFeed.Channel.Title)
	fmt.Printf("| %v\n", rssFeed.Channel.Description)
	fmt.Println("|")
	fmt.Println("|")

	for _, channel := range rssFeed.Channel.Item {
		fmt.Printf("| %v\n", channel.Title)
	}

	fmt.Println("|")
	fmt.Println("---------------------")

	return nil
}
