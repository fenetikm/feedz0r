package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/commands"
	"github.com/fenetikm/feedz0r/internal/commands/add"
	"github.com/fenetikm/feedz0r/internal/commands/fetch"
	"github.com/fenetikm/feedz0r/internal/commands/help"
	"github.com/fenetikm/feedz0r/internal/commands/list"
	"github.com/fenetikm/feedz0r/internal/commands/post"
	"github.com/fenetikm/feedz0r/internal/commands/watch"
	"github.com/fenetikm/feedz0r/internal/config"
	"github.com/fenetikm/feedz0r/internal/db/database"
	"github.com/fenetikm/feedz0r/internal/state"
	_ "modernc.org/sqlite"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	cmds := commands.Commands{
		Handlers: make(map[string]func(*state.State, cmdtypes.Command) error),
	}

	db, err := sql.Open("sqlite", "feedz0r.db")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Can't open db")
	}

	dbQueries := database.New(db)
	defer db.Close()

	s := state.State{
		Config: &c,
		Cmds:   &cmds,
		Db:     dbQueries,
	}

	s.Cmds.Register("help", help.Handle)
	s.Cmds.Register("add", add.Handle)
	s.Cmds.Register("fetch", fetch.Handle)
	s.Cmds.Register("watch", watch.Handle)
	s.Cmds.Register("list", list.Handle)
	s.Cmds.Register("post", post.Handle)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough args")
		os.Exit(1)
	}

	argCmd := cmdtypes.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = s.Cmds.Run(&s, argCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
