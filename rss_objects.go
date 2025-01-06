package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
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
