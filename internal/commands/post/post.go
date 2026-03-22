package post

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Missing required arg: <post id>")
	}

	id, err := strconv.ParseInt(cmd.Args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("Invalid post id: %w", err)
	}

	post, err := s.Db.GetPostByID(context.Background(), id)
	if err != nil {
		return fmt.Errorf("Couldn't get post by ID: %w", err)
	}

	fmt.Println(post.Description)

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	return nil
}
