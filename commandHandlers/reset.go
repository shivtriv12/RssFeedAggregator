package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func ResetHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	err := s.Db.ResetUserTable(ctx)
	if err != nil {
		fmt.Println("Error in deleting everything from users table")
		os.Exit(1)
	}
	fmt.Println("Deleted all rows from users table")
	return nil
}
