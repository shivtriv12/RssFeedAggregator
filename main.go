package main

import (
	"fmt"
	"os"

	commands "github.com/shivtriv12/BlogAggregator/commandHandlers"
	"github.com/shivtriv12/BlogAggregator/internal/config"
	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func main() {
	currentConfig, err := config.Read()
	if err != nil {
		fmt.Println("Error Reading gaterConfig")
	}
	currentConfigState := types.State{
		ConfigState: currentConfig,
	}
	fmt.Println(currentConfigState.ConfigState.Current_User_Name, currentConfigState.ConfigState.Db_Url)
	mapCommands := types.Commands{
		CommandsMap: make(map[string]func(s *types.State, cmd types.Command) error),
	}
	mapCommands.Register("login", commands.LoginHandler)
	cmd := os.Args[1:]
	if len(cmd) < 2 {
		fmt.Println("not enough commands")
		os.Exit(1)
	}
	userCmd := types.Command{
		Name: cmd[0],
		Args: cmd[1:],
	}
	if err := mapCommands.Run(&currentConfigState, userCmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
