package types

import (
	"errors"

	"github.com/shivtriv12/BlogAggregator/internal/config"
)

type State struct {
	ConfigState *config.Config `json:"configState"`
}

type Command struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
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
