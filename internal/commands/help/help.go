package help

import (
	"fmt"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("Help!")

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("help - prints out help")

	return nil
}
