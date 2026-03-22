package fetch

import (
	"context"
	"fmt"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/feeds"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	rssFeed, err := feeds.Process(s, context.Background())
	if err != nil {
		fmt.Println("Error fetching a feed.")
		return err
	}

	fmt.Printf("Fetched feed %s", rssFeed.Channel.Title)

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("fetch - fetch a feed")

	return nil
}
