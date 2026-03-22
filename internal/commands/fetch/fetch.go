package fetch

import (
	"context"
	"errors"
	"fmt"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/feeds"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return errors.New("Error fetching next feed.")
	}

	fmt.Printf("Fetching feed: %s\n", feed.Url)

	rssFeed, err := feeds.Fetch(s, context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed: %v", err)
		return nil
	}

	err = s.Db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("Error marking feed as fetched.")
	}

	fmt.Println(rssFeed.Channel.Title)

	// for _, item := range rssFeed.Channel.Item {
	// }

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("fetch - fetch a feed")

	return nil
}
