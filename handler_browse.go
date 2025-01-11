package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2

	if len(cmd.Args) == 1 {
		arg, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("could not parse argument: %v", err)
		}
		limit = arg
	}

	args := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.database.GetPostsForUser(context.Background(), args)
	if err != nil {
		return fmt.Errorf("error getting posts for user")
	}

	for _, post := range posts {
		fmt.Println()
		fmt.Println()
		fmt.Println("---------------------")
		fmt.Println("|")
		fmt.Printf("| %v\n", post.Title)
		fmt.Printf("| %v\n", post.PublishedAt)
		fmt.Println("|")
		fmt.Println("|")
		fmt.Printf("| %v\n", post.Description)
		fmt.Println("|")
		fmt.Println("---------------------")
	}

	return nil
}
