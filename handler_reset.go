package main

import (
	"context"
	"fmt"
)

func HandlerReset(s *state, cmd command) error {
	err := s.database.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("could not truncate users table: %v", err)
	}

	fmt.Println("Users table successfully truncated")

	return nil
}
