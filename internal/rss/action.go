package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	var feeds RSSFeed
	if err := xml.Unmarshal(body, &feeds); err != nil {
		return nil, fmt.Errorf("error parsing XML: %w", err)
	}
	feeds.Channel.Description = html.UnescapeString(feeds.Channel.Description)
	feeds.Channel.Title = html.UnescapeString(feeds.Channel.Title)
	for i, feed := range feeds.Channel.Item {
		feeds.Channel.Item[i].Description = html.UnescapeString(feed.Description)
		feeds.Channel.Item[i].Title = html.UnescapeString(feed.Title)
	}
	return &feeds, nil
}
