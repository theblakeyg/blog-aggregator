package main

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *state, cmd command) error {
	result, err := s.database.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get all users: %v", err)
	}

	for _, feed := range result {
		name := feed.Name.String
		url := feed.Url.String
		user, err := s.database.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("could not get user by user_id: %v", err)
		}

		fmt.Printf("Name: %v | URL: %v | UserName: %v\n", name, url, user.Name.String)
	}

	return nil
}
