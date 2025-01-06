package main

import (
	"context"
	"database/sql"
	"fmt"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	userName := cmd.Args[0]

	user, err := s.database.GetUser(context.Background(), sql.NullString{String: userName, Valid: true})
	if err != nil {
		return fmt.Errorf("could not get user by this username: %v", err)
	}

	err = s.config.SetUser(user.Name.String)
	if err != nil {
		return fmt.Errorf("could not set current user: %v", err)
	}

	fmt.Printf("User has been set to: %v", user.Name.String)

	return nil
}
