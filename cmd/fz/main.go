package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fenetikm/feedz0r/internal/commands"
	"github.com/fenetikm/feedz0r/internal/config"
	"github.com/fenetikm/feedz0r/internal/state"
)

// todo: shift this to module
func handlerHelp(s *state.State, cmd commands.Command) error {
	fmt.Println("Help!")

	return nil
}

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := state.State{
		Config: &c,
	}

	cmds := commands.Commands{
		Handlers: make(map[string]func(*state.State, commands.Command) error),
	}
	cmds.Register("help", handlerHelp)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough args")
		os.Exit(1)
	}

	argCmd := commands.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.Run(&s, argCmd)
	if err != nil {
		fmt.Printf("Command %s does not exist.\n", argCmd.Name)
		os.Exit(1)
	}
}
