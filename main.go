package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	commands "github.com/shivtriv12/BlogAggregator/commandHandlers"
	"github.com/shivtriv12/BlogAggregator/internal/config"
	"github.com/shivtriv12/BlogAggregator/internal/database"
	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func main() {
	//db setup
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/gator")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	defer db.Close()
	dbQueries := database.New(db)
	// getting current config
	currentConfig, err := config.Read()
	if err != nil {
		fmt.Println("Error Reading gaterConfig")
	}
	//setting current config and db in state
	appState := types.State{
		Db:          dbQueries,
		ConfigState: currentConfig,
	}
	// commands registerations
	mapCommands := types.Commands{
		CommandsMap: make(map[string]func(s *types.State, cmd types.Command) error),
	}
	mapCommands.Register("login", commands.LoginHandler)
	mapCommands.Register("register", commands.RegisterHandler)
	mapCommands.Register("reset", commands.ResetHandler)
	mapCommands.Register("users", commands.UsersHandler)
	mapCommands.Register("agg", commands.AggHandler)
	mapCommands.Register("addfeed", commands.AddFeedHandler)
	// user commands
	cmd := os.Args[1:]
	if len(cmd) < 1 {
		fmt.Println("not enough commands")
		os.Exit(1)
	}
	userCmd := types.Command{
		Name: cmd[0],
		Args: cmd[1:],
	}
	if err := mapCommands.Run(&appState, userCmd); err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
}
