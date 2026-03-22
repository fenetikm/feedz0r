package add

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/db/database"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	// todo: params validation
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Missing two required args: <name> <url>")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	currentTime := time.Now().Unix()

	createParams := database.CreateFeedParams{
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      sql.NullString{String: cmd.Args[0], Valid: true},
		Url:       cmd.Args[1],
	}

	_, err := s.Db.CreateFeed(ctx, createParams)
	if err != nil {
		return errors.New("Couldn't create feed.")
	}

	fmt.Printf("Feed %s added.", cmd.Args[1])

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("add - add a feed by name and URL\n")
	fmt.Println("Usage: fz add <name> <url>\n")
	fmt.Println("Example:")
	fmt.Println(`  fz add "its.mw" https://its.mw/atom.xml`)

	return nil
}
