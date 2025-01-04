package main

import (
	"fmt"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	userName := cmd.Args[0]

	err := s.config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("could not set current user: %v", err)
	}

	fmt.Printf("User has been set to: %v", userName)

	return nil
}
