package main

import (
	"context"
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
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	result, err := s.database.CreateUser(context.Background(), args)
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}

	fmt.Printf("User has been added: %v", result.Name)

	err = s.config.SetUser(result.Name)
	if err != nil {
		return fmt.Errorf("could not login with new user: %v", err)
	}

	return nil
}
