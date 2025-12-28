package commands

import (
	"context"
	"fmt"

	"github.com/shivtriv12/RSSFeedAggregator/internal/config"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func UsersHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	users, err := s.Db.GetAllUsers(ctx)
	if err != nil {
		return err
	}
	currentUser, err := config.Read()
	if err != nil {
		return fmt.Errorf("Unable to get current user %w", err)
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
