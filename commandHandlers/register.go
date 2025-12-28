package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shivtriv12/RSSFeedAggregator/internal/config"
	"github.com/shivtriv12/RSSFeedAggregator/internal/database"
	"github.com/shivtriv12/RSSFeedAggregator/internal/types"
)

func RegisterHandler(s *types.State, cmd types.Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("the register handler expects a single argument, the username")
	}
	ctx := context.Background()
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}
	user, err := s.Db.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("Error creating user: %w", err)
	}
	s.ConfigState.Current_User_Name = cmd.Args[0]
	err = config.SetUser(s.ConfigState)
	if err != nil {
		return fmt.Errorf("Error saving gatorconfig: %w", err)
	}
	fmt.Println("user was created", user.Name)
	return nil
}
