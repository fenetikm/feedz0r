package commands

import (
	"errors"

	"github.com/fenetikm/feedz0r/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	if fn, ok := c.Handlers[cmd.Name]; ok {
		return fn(s, cmd)
	}

	return errors.ErrUnsupported
}
