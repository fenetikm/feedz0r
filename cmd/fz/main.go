package main

import (
	"fmt"
	"log"

	"github.com/fenetikm/feedz0r/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
