package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

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
