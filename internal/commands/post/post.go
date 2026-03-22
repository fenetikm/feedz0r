package post

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"time"

	glamour "charm.land/glamour/v2"
	htmltomd "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/state"
)

func Handle(s *state.State, cmd cmdtypes.Command) error {
	fs := flag.NewFlagSet("post", flag.ContinueOnError)
	outputFormat := fs.String("output", "basic", "Output format: basic or markdown")
	if err := fs.Parse(cmd.Args); err != nil {
		return err
	}

	if len(fs.Args()) < 1 {
		return fmt.Errorf("Missing required arg: <post id>")
	}

	id, err := strconv.ParseInt(fs.Args()[0], 10, 64)
	if err != nil {
		return fmt.Errorf("Invalid post id: %w", err)
	}

	post, err := s.Db.GetPostByID(context.Background(), id)
	if err != nil {
		return fmt.Errorf("Couldn't get post by ID: %w", err)
	}

	switch *outputFormat {
	case "markdown":
		t := time.Unix(post.PublishedAt, 0)
		date := t.Format("Mon, 02 Jan. 2006")
		content := fmt.Sprintf(`<h2>%s</h2>
<a href="%s">%s</a><p>%s</p>%s`, post.Title, post.Url, post.Url, date, post.Description.String)
		md, err := htmltomd.ConvertString(content)
		if err != nil {
			return fmt.Errorf("Couldn't convert description to markdown: %w", err)
		}

		out, err := glamour.Render(md, "dark")
		fmt.Print(out)

		return err
	default:
		fmt.Println(post.Title)
		t := time.Unix(post.PublishedAt, 0)
		fmt.Println(t.Format("Mon, 02 Jan. 2006"))
		fmt.Println(post.Description)
	}

	return nil
}

func Help(s *state.State, cmd cmdtypes.Command) error {
	return nil
}
