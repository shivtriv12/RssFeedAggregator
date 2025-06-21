package commands

import (
	"context"
	"fmt"

	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func FeedsHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	feeds, err := s.Db.GetAllFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error getting stored feeds %w", err)
	}
	for i, feed := range feeds {
		fmt.Println(i + 1)
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserName)
	}
	return nil
}
