package commands

import (
	"context"
	"fmt"

	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func FollowingHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	user, err := s.Db.GetUser(ctx, s.ConfigState.Current_User_Name)
	if err != nil {
		return fmt.Errorf("error in getting current user %w", err)
	}
	feed_follows, err := s.Db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error fetching followed feeds: %w", err)
	}
	if len(feed_follows) == 0 {
		fmt.Printf("User '%s' is not following any feeds\n", user.Name)
		return nil
	}
	for i, follow := range feed_follows {
		fmt.Printf("%d. %s\n", i+1, follow.FeedName)
	}
	return nil
}
