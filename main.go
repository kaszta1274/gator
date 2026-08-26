package main

import (
	"log"
	"os"

	"github.com/kaszta1274/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = cmds.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
