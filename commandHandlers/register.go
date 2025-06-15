package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/shivtriv12/BlogAggregator/internal/config"
	"github.com/shivtriv12/BlogAggregator/internal/database"
	"github.com/shivtriv12/BlogAggregator/internal/types"
)

func RegisterHandler(s *types.State, cmd types.Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
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
		fmt.Println("Error creating user:", err)
		os.Exit(1)
	}
	s.ConfigState.Current_User_Name = cmd.Args[0]
	err = config.SetUser(s.ConfigState)
	if err != nil {
		fmt.Println("Error saving config:", err)
		os.Exit(1)
	}
	fmt.Println("user was created", user.Name)
	return nil
}
