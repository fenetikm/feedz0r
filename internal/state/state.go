package state

import (
	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/config"
)

type Runner interface {
	Register(name string, f func(*State, cmdtypes.Command) error)
	Run(s *State, cmd cmdtypes.Command) error
}

type State struct {
	Config *config.Config
	Cmds   Runner
}
