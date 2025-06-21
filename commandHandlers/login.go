package commands

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/shivtriv12/RSSFeedAggregator/internal/config"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func LoginHandler(s *types.State, cmd types.Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("the login handler expects a single argument, the username")
	}
	ctx := context.Background()
	user, err := s.Db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		os.Exit(1)
	}
	s.ConfigState.Current_User_Name = cmd.Args[0]
	err = config.SetUser(s.ConfigState)
	if err != nil {
		fmt.Println("Error updating config:", err)
		os.Exit(1)
	}
	fmt.Println(user.Name)
	return nil
}
