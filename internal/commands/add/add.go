package add

import (
	"fmt"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("Add")
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Missing URL.")
	}

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("add - add a feed")

	return nil
}
