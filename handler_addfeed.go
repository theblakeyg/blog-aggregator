package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("not enough arguments provided")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]
	userName := s.config.CurrentUserName
	user, err := s.database.GetUser(context.Background(), sql.NullString{String: userName, Valid: true})
	if err != nil {
		return fmt.Errorf("error getting user id: %v", err)
	}

	args := database.CreateFeedParams{
		ID:        uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      sql.NullString{String: name, Valid: true},
		Url:       sql.NullString{String: url, Valid: true},
		UserID:    uuid.NullUUID{UUID: user.ID.UUID, Valid: true},
	}

	result, err := s.database.CreateFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error creating feed: %v", err)
	}

	fmt.Printf("added RSS feed successfully: %v\n", result)
	return nil
}
