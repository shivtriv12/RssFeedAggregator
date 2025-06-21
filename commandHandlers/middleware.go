package commands

import (
	"context"

	"github.com/shivtriv12/RSSFeedAggregator/internal/database"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func MiddlewareLoggedIn(handler func(s *types.State, cmd types.Command, user database.User) error) func(*types.State, types.Command) error {
	return func(s *types.State, cmd types.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.ConfigState.Current_User_Name)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
