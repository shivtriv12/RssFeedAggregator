package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func ResetHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	err := s.Db.ResetUserTable(ctx)
	if err != nil {
		fmt.Println("Error in resetting")
		os.Exit(1)
	}
	fmt.Println("DB resetted")
	return nil
}
