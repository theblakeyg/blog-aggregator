package main

import (
	"fmt"
	"log"
	"os"

	"github.com/theblakeyg/blog-aggregator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	//Read our config file
	configFile, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v", err)
	}

	//Create our current state and attach our config file
	currentState := &state{
		config: &configFile,
	}

	//Register all of our commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.Register("login", HandlerLogin)

	//Check to see that we have enough arguments
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments provided")
		return
	}

	//Separate cmdName and cmdArgs from all arguments
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	//Run the provided command with the provided arguments and the current state
	err = cmds.Run(currentState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
