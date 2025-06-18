package commands

import (
	"context"
	"fmt"

	"github.com/shivtriv12/BlogAggregator/internal/rss"
	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func AggHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"

	fmt.Println("📰 Fetching RSS feed from:", url)
	feeds, err := rss.FetchFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	fmt.Println("\n============ FEED INFORMATION ============")
	fmt.Printf("📑 Title: %s\n", feeds.Channel.Title)
	fmt.Printf("🔗 Link: %s\n", feeds.Channel.Link)
	fmt.Printf("📝 Description: %s\n", feeds.Channel.Description)

	fmt.Printf("\n============ ARTICLES (%d) ============\n", len(feeds.Channel.Item))

	for i, feed := range feeds.Channel.Item {
		fmt.Printf("\n--- Article %d ---\n", i+1)
		fmt.Printf("📌 Title: %s\n", feed.Title)
		fmt.Printf("🕒 Published: %s\n", feed.PubDate)
		fmt.Printf("🔗 Link: %s\n", feed.Link)
		fmt.Printf("📄 Summary: %s\n", feed.Description)
	}

	return nil
}
