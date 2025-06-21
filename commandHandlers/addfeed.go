package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shivtriv12/RSSFeedAggregator/internal/database"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func AddFeedHandler(s *types.State, cmd types.Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed requires exactly 2 arguments: name and url")
	}
	ctx := context.Background()
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}
	feed, err := s.Db.CreateFeed(ctx, params)
	if err != nil {
		return fmt.Errorf("error adding feed %w", err)
	}
	fmt.Println("Feed added successfully!")
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("Owner: %s (ID: %s)\n", s.ConfigState.Current_User_Name, feed.UserID)
	fmt.Printf("Created at: %s\n", feed.CreatedAt)

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.Db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("feed created but failed to follow it: %w", err)
	}
	fmt.Printf("\nAutomatically followed feed: %s\n", feedFollow.FeedName)

	return nil
}
