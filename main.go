package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/theblakeyg/blog-aggregator/internal/config"
	"github.com/theblakeyg/blog-aggregator/internal/database"
)

type state struct {
	config   *config.Config
	database *database.Queries
}

func main() {
	//Read our config file
	configFile, err := config.Read()
	if err != nil {
		log.Fatal("error reading config:", err)
	}

	//Connect to our db
	db, err := sql.Open("postgres", configFile.DbUrl)
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	dbQueries := database.New(db)

	//Create our current state and attach our config file
	currentState := &state{
		config:   &configFile,
		database: dbQueries,
	}

	//Register all of our commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.Register("login", HandlerLogin)
	cmds.Register("register", HandlerRegister)

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
