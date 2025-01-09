package main

import (
	"context"
	"fmt"

	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandlerFollowing(s *state, cmd command, user database.User) error {
	// user, err := s.database.GetUser(context.Background(), sql.NullString{String: s.config.CurrentUserName, Valid: true})
	// if err != nil {
	// 	return fmt.Errorf("error getting current user: %v", err)
	// }

	feeds, err := s.database.FollowsByUserId(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting follow by userid: %v", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
	}

	return nil
}
