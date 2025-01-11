package main

import (
	"fmt"
	"time"
)

func HandlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enough arguments provided")
	}

	arg := cmd.Args[0]

	duration, err := time.ParseDuration(arg)
	if err != nil {
		return fmt.Errorf("error parsing duration: %v", err)
	}
	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
