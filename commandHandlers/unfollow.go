package commands

import (
	"context"
	"fmt"

	"github.com/shivtriv12/BlogAggregator/internal/database"
	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func UnfollowHandler(s *types.State, cmd types.Command, user database.User) error {
	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.Args[0],
	}
	err := s.Db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println("Deleted feedfollow successfully")
	return nil
}
