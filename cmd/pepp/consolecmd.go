package main

import (
	"os"

	"github.com/pepperdb/pepperdb-core/cmd/console"
	"github.com/urfave/cli"
)

var (
	consoleCommand = cli.Command{
		Action:   MergeFlags(consoleStart),
		Name:     "console",
		Usage:    "Start an interactive JavaScript console",
		Category: "CONSOLE COMMANDS",
		Description: `
The Neb console is an interactive shell for the JavaScript runtime environment.`,
	}
)

func consoleStart(ctx *cli.Context) error {
	neb, err := makePepp(ctx)
	if err != nil {
		return err
	}

	console := console.New(console.Config{
		Prompter:   console.Stdin,
		PrompterCh: make(chan string),
		Writer:     os.Stdout,
		Neb:        neb,
	})

	console.Setup()
	console.Interactive()
	defer console.Stop()
	return nil
}
