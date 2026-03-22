package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fenetikm/feedz0r/internal/cmdtypes"
	"github.com/fenetikm/feedz0r/internal/commands"
	"github.com/fenetikm/feedz0r/internal/commands/help"
	"github.com/fenetikm/feedz0r/internal/config"
	"github.com/fenetikm/feedz0r/internal/state"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	cmds := commands.Commands{
		Handlers: make(map[string]func(*state.State, cmdtypes.Command) error),
	}

	s := state.State{
		Config: &c,
		Cmds:   &cmds,
	}

	s.Cmds.Register("help", help.Handle)

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
		fmt.Printf("Command %s does not exist.\n", argCmd.Name)
		os.Exit(1)
	}
}
