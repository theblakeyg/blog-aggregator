package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) Run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	return f(s, cmd)
}
