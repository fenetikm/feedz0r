package commands

import (
	"errors"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/state"
)

type Commands struct {
	Handlers map[string]func(*state.State, cmdtypes.Command) error
}

// Implement Runner.Register
func (c *Commands) Register(name string, f func(*state.State, cmdtypes.Command) error) {
	c.Handlers[name] = f
}

// Implement Runner.Run
func (c *Commands) Run(s *state.State, cmd cmdtypes.Command) error {
	if fn, ok := c.Handlers[cmd.Name]; ok {
		return fn(s, cmd)
	}

	return errors.ErrUnsupported
}
