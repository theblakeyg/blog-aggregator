package main

import (
	"context"
	"fmt"

	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	feed, err := s.database.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error getting feed by url: %v", err)
	}

	args := database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.database.UnfollowFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %v", err)
	}

	fmt.Println("feed successfully unfollowed")

	return nil
}
