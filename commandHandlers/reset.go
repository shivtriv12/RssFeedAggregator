package commands

import (
	"context"
	"fmt"

	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func ResetHandler(s *types.State, cmd types.Command) error {
	ctx := context.Background()
	err := s.Db.ResetUserTable(ctx)
	if err != nil {
		return err
	}
	err = s.Db.ResetFeedsTable(ctx)
	if err != nil {
		return err
	}
	fmt.Println("DB resetted")
	return nil
}
