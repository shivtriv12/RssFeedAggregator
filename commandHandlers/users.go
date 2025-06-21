package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/shivtriv12/RSSFeedAggregator/internal/config"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func UsersHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	users, err := s.Db.GetAllUsers(ctx)
	if err != nil {
		fmt.Println("get all users query didnt ran")
		os.Exit(1)
	}
	currentUser, err := config.Read()
	if err != nil {
		fmt.Println("Unable to get current user")
		os.Exit(1)
	}
	for _, user := range users {
		if user.Name == currentUser.Current_User_Name {
			fmt.Println("* " + user.Name + " (current)")
			continue
		}
		fmt.Println("* " + user.Name)
	}
	return nil
}
