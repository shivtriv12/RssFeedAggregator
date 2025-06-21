package commands

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shivtriv12/RSSFeedAggregator/internal/database"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func FollowHandler(s *types.State, cmd types.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow requires exactly 1 argument: feed URL")
	}
	ctx := context.Background()
	feed, err := s.Db.GetFeedByUrl(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("feed not found with URL %s - please add it first using 'addfeed'", cmd.Args[0])
	}
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	follow, err := s.Db.CreateFeedFollow(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return fmt.Errorf("you're already following this feed")
		}
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("User '%s' is now following feed '%s'\n", follow.UserName, follow.FeedName)
	return nil
}
