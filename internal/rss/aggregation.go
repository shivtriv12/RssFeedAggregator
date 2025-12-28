package rss

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shivtriv12/RSSFeedAggregator/internal/database"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
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

	fmt.Printf("Found %d items in feed\n", len(feedContent.Channel.Item))

	for _, item := range feedContent.Channel.Item {
		publishedAt, err := parsePublishedDate(item.PubDate)
		if err != nil {
			fmt.Printf("Error parsing date for post %s: %v\n", item.Title, err)
			publishedAt = time.Now()
		}

		_, err = s.Db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			fmt.Printf("Error saving post %s: %v\n", item.Title, err)
		} else {
			fmt.Printf("saved post: %s\n", item.Title)
		}
	}

	err = s.Db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	return nil
}

func parsePublishedDate(dateStr string) (time.Time, error) {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822,
		time.RFC822Z,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}
