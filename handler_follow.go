package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	feedUrl := cmd.Args[0]

	user, err := s.database.GetUser(context.Background(), sql.NullString{String: s.config.CurrentUserName, Valid: true})
	if err != nil {
		return fmt.Errorf("could not get current user from config: %v", err)
	}

	feed, err := s.database.GetFeedByUrl(context.Background(), sql.NullString{String: feedUrl, Valid: true})
	if err != nil {
		return fmt.Errorf("could not get feed by url: %v", err)
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UserID:    uuid.NullUUID{UUID: user.ID.UUID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feed.ID.UUID, Valid: true},
	}

	result, err := s.database.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Printf("User (%v) is now following Feed (%v)", result.UserName.String, result.FeedName.String)

	return nil
}

func HandlerFollowing(s *state, cmd command) error {
	user, err := s.database.GetUser(context.Background(), sql.NullString{String: s.config.CurrentUserName, Valid: true})
	if err != nil {
		return fmt.Errorf("error getting current user: %v", err)
	}

	feeds, err := s.database.FollowsByUserId(context.Background(), uuid.NullUUID{UUID: user.ID.UUID, Valid: true})
	if err != nil {
		return fmt.Errorf("error getting follow by userid: %v", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
	}

	return nil
}
