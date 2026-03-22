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

	fmt.Println(feed.Url)

	rssFeed, err := feeds.Fetch(s, context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("Error fetching feed: %v", err)
		return nil
	}

	fmt.Println(rssFeed.Channel.Link)
	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item)
	}

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("fetch - fetch a feed")

	return nil
}
