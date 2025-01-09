package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	feedUrl := cmd.Args[0]

	feed, err := s.database.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("could not get feed by url: %v", err)
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	result, err := s.database.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Printf("User (%v) is now following Feed (%v)", result.UserName, result.FeedName)

	return nil
}
