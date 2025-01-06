package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theblakeyg/blog-aggregator/internal/database"
)

func HandlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	userName := cmd.Args[0]

	args := database.CreateUserParams{
		ID:        uuid.NullUUID{UUID: uuid.New(), Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      sql.NullString{String: userName, Valid: true},
	}

	result, err := s.database.CreateUser(context.Background(), args)
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}

	fmt.Printf("User has been added: %v", result.Name.String)

	err = s.config.SetUser(result.Name.String)
	if err != nil {
		return fmt.Errorf("could not login with new user: %v", err)
	}

	return nil
}
