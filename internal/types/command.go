package types

import (
	"errors"

	"github.com/shivtriv12/RSSFeedAggregator/internal/config"
	"github.com/shivtriv12/RSSFeedAggregator/internal/database"
)

type State struct {
	ConfigState *config.Config
	Db          *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandsMap map[string]func(s *State, cmd Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.CommandsMap[cmd.Name]
	if !ok {
		return errors.New("no command handler found")
	}
	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) error {
	if c.CommandsMap == nil {
		c.CommandsMap = make(map[string]func(s *State, cmd Command) error)
	}
	c.CommandsMap[name] = f
	return nil
}
