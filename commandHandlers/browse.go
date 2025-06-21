package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/shivtriv12/BlogAggregator/internal/database"
	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func BrowseHandler(s *types.State, cmd types.Command, user database.User) error {
	limit := int32(2)

	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit value: %s", cmd.Args[0])
		}
		limit = int32(parsedLimit)
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found. Try following some feeds first.")
		return nil
	}

	fmt.Printf("Latest %d posts:\n", len(posts))
	for i, post := range posts {
		fmt.Printf("\n=== Post %d ===\n", i+1)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Published: %s\n", post.PublishedAt.Format("Jan 02, 2006"))
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description)
	}

	return nil
}
