package rss

import (
	"context"
	"fmt"

	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func ScrapeFeeds(s *types.State) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get next feed: %w", err)
	}

	fmt.Printf("Fetching feed: %s (%s)\n", feed.Name, feed.Url)

	feedContent, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed %s: %w", feed.Url, err)
	}

	err = s.Db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	fmt.Printf("Found %d items in feed\n", len(feedContent.Channel.Item))
	for i, item := range feedContent.Channel.Item {
		fmt.Printf("%d. %s\n", i+1, item.Title)
	}
	return nil
}
