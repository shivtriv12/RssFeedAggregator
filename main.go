package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	commands "github.com/shivtriv12/RSSFeedAggregator/commandHandlers"
	"github.com/shivtriv12/RSSFeedAggregator/internal/config"
	"github.com/shivtriv12/RSSFeedAggregator/internal/database"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env: ", err)
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL not set: ", err)
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	currentConfig, err := config.Read()
	if err != nil {
		log.Fatal("Error Reading gaterConfig: ", err)
	}
	appState := types.State{
		Db:          dbQueries,
		ConfigState: currentConfig,
	}

	mapCommands := types.Commands{
		CommandsMap: make(map[string]func(s *types.State, cmd types.Command) error),
	}
	mapCommands.Register("register", commands.RegisterHandler)
	mapCommands.Register("login", commands.LoginHandler)
	mapCommands.Register("users", commands.UsersHandler)
	mapCommands.Register("reset", commands.ResetHandler)
	mapCommands.Register("addfeed", commands.MiddlewareLoggedIn(commands.AddFeedHandler))
	mapCommands.Register("feeds", commands.FeedsHandler)
	mapCommands.Register("follow", commands.MiddlewareLoggedIn(commands.FollowHandler))
	mapCommands.Register("following", commands.MiddlewareLoggedIn(commands.FollowingHandler))
	mapCommands.Register("unfollow", commands.MiddlewareLoggedIn(commands.UnfollowHandler))
	mapCommands.Register("agg", commands.AggHandler)
	mapCommands.Register("browse", commands.MiddlewareLoggedIn(commands.BrowseHandler))

	cmd := os.Args[1:]
	if len(cmd) < 1 {
		log.Fatal("Usage: go run . <cmd> <args>")
	}
	userCmd := types.Command{
		Name: cmd[0],
		Args: cmd[1:],
	}
	if err := mapCommands.Run(&appState, userCmd); err != nil {
		log.Fatal("Error running command: ", err)
	}
}
