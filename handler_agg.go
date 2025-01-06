package main

import (
	"context"
	"fmt"
)

func HandlerAgg(s *state, cmd command) error {
	url := "http://www.wagslane.dev/index.xml"
	rss, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Println(rss)

	return nil

}
