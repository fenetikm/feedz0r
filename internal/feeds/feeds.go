package feeds

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/fenetikm/feedz0r/internal/state"
)

// todo: Atom support, and generic "feed" struct
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func Fetch(s *state.State, ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(ctx, time.Duration(s.Config.Fetch.Timeout)*time.Second)
	defer cancel()

	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error creating request: %v\n", err)
	}

	req.Header.Set("User-Agent", "feedz0r")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error making request: %v\n", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error creating request: %v\n", err)
	}

	var feed RSSFeed
	err = xml.Unmarshal([]byte(data), &feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error unmarshalling xml: %v\n", err)
	}

	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
	}

	return &feed, nil
}
