package commands

import (
	"errors"

	"github.com/shivtriv12/BlogAggregator/internal/config"
	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func LoginHandler(s *types.State, cmd types.Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	s.ConfigState.Current_User_Name = cmd.Args[0]
	return config.SetUser(s.ConfigState)
}
