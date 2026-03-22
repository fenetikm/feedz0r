package watch

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/feeds"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	mins := fmt.Sprintf("%dm", s.Config.Fetch.RefreshMins)

	timeBetween, err := time.ParseDuration(mins)
	if err != nil {
		log.Fatalln("Could not parse time duration argument.")
	}
	fmt.Printf("Watching feeds and processing every %s\n", timeBetween)

	ticker := time.NewTicker(timeBetween)
	for ; ; <-ticker.C {
		rssFeed, err := feeds.Process(s, context.Background())
		if err != nil {
			log.Fatalln("Could not scrape feeds.")
		}
		fmt.Printf("Processed feed: %s %s", rssFeed.Channel.Title, rssFeed.Channel.Link)
		fmt.Println("\n...waiting for next run...\n")
	}
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	fmt.Println("watch feeds")

	return nil
}
