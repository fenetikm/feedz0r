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
	limit := fs.Int("limit", 20, "Limit number of items returned")
	listType := fs.String("type", "feeds", "Item type: feeds or posts")
	fs.Parse(cmd.Args)

	switch *listType {
	case "posts":
		posts, err := s.Db.GetPosts(context.Background(), int64(*limit))
		if err != nil {
			fmt.Println("Error, could not get any posts.")
			return nil
		}
		for _, post := range posts {
			fmt.Printf("%s (%d) %s\n", post.Title, post.ID, post.Url)
		}
	default:
		feeds, err := s.Db.GetFeeds(context.Background(), int64(*limit))
		if err != nil {
			fmt.Println("Error, could not get any feeds.")
			return nil
		}
		for _, feed := range feeds {
			fmt.Printf("%s (%d) %s\n", feed.Name.String, feed.ID, feed.Url)
		}
	}

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	return nil
}
