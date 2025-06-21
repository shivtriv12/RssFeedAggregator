package commands

import (
	"fmt"
	"time"

	"github.com/shivtriv12/RSSFeedAggregator/internal/rss"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func AggHandler(s *types.State, cmd types.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("excepts one argument time between requests")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		fmt.Println("unable to parse time passed")
	}
	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("fetching next feed")
		rss.ScrapeFeeds(s)
	}
}
