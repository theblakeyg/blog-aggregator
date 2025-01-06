package main

import (
	"context"
	"fmt"
)

func HandlerUsers(s *state, cmd command) error {
	result, err := s.database.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get all users: %v", err)
	}

	for _, user := range result {
		name := user.Name.String
		if name == s.config.CurrentUserName {
			fmt.Printf("* %v (current)\n", name)
		} else {
			fmt.Printf("* %v\n", name)
		}
	}

	return nil
}
