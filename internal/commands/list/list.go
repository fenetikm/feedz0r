package list

import (
	"context"
	"flag"
	"fmt"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	limit := fs.Int("limit", 0, "Limit number of items returned (0 = no limit)")
	fs.Parse(cmd.Args)
	listType := "feeds"

	// todo:
	// - feeds vs posts

	if listType == "feeds" {
		feeds, err := s.Db.GetFeeds(context.Background(), int64(*limit))
		if err != nil {
			fmt.Println("Error, could not get any feeds.")
			return nil
		}

		for _, feed := range feeds {
			fmt.Printf("%s (%d) %s\n", feed.Name.String, feed.ID, feed.Url)
		}

		return nil
	}

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	return nil
}
