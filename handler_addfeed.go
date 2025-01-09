package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("not enough arguments provided")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	args := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	result, err := s.database.CreateFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Printf("added RSS feed successfully: %v\n", result)

	err = HandlerFollow(s, command{Args: []string{result.Url}}, user)
	if err != nil {
		return fmt.Errorf("error following newly created feed: %v", err)
	}

	fmt.Println("RSS feed successfully followed")

	return nil
}
