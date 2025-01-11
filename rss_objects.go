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

	"github.com/google/uuid"
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

	for _, item := range rssFeed.Channel.Item {
		pubDate, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			fmt.Printf("could not parse published data: %v\n", err)
			return fmt.Errorf("could not parse published data: %v", err)
		}

		args := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}
		result, err := s.database.CreatePost(context.Background(), args)
		if err != nil {
			fmt.Printf("error creating post: %v\n", err)

			return fmt.Errorf("error creating post: %v", err)
		}

		fmt.Printf("Post '%v' from Feed '%v' saved succesfully\n", result.Title, feed.Name)
	}

	return nil
}
