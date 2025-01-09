package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	feed, err := s.database.GetFeedByUrl(context.Background(), sql.NullString{String: cmd.Args[0], Valid: true})
	if err != nil {
		return fmt.Errorf("error getting feed by url: %v", err)
	}

	args := database.UnfollowFeedParams{
		UserID: uuid.NullUUID{UUID: user.ID.UUID, Valid: true},
		FeedID: uuid.NullUUID{UUID: feed.ID.UUID, Valid: true},
	}

	err = s.database.UnfollowFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}

	fmt.Println("feed successfully unfollowed")

	return nil
}
