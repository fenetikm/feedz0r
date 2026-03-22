package feeds

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/fenetikm/feedz0r/internal/db/database"
	"github.com/fenetikm/feedz0r/internal/state"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
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

func Process(s *state.State, ctx context.Context) (*RSSFeed, error) {
	// todo: output how many new posts
	feed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return &RSSFeed{}, errors.New("Error fetching next feed.")
	}

	rssFeed, err := Fetch(s, ctx, feed.Url)
	if err != nil {
		return &RSSFeed{}, err
	}

	err = s.Db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error marking feed as fetched: %v\n", err)
	}

	currentTime := time.Now().Unix()
	for _, item := range rssFeed.Channel.Item {
		pubDate, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", item.PubDate)
		if err != nil {
			fmt.Printf("Could not convert published date %s.\n", item.PubDate)
			continue
		}
		postParams := database.CreatePostParams{
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: pubDate.Unix(),
			FeedID:      feed.ID,
		}
		_, err = s.Db.CreatePost(ctx, postParams)
		if err != nil {
			var sqliteErr *sqlite.Error
			if errors.As(err, &sqliteErr) && sqliteErr.Code() ==
				sqlite3.SQLITE_CONSTRAINT_UNIQUE {
				continue // skip duplicates
			}

			return &RSSFeed{}, fmt.Errorf("Error storing post: %v\n", err)
		}
	}

	return rssFeed, nil
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
