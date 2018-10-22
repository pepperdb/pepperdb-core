package main

import (
	"github.com/pepperdb/pepperdb-core/dappserver"
	"github.com/urfave/cli"
)

var (
	dappServerCommand = cli.Command{
		Action:    _dappServer,
		Name:      "dappserver",
		Usage:     "Start dapp server",
		ArgsUsage: "<dappserverconfig>",
		Category:  "DAPPSERVER COMMANDS",
		Description: `
The dappserver command start a dapp server.`,
	}
)

func _dappServer(ctx *cli.Context) error {
	ds, err := dappserver.NewDAppServer()
	if err != nil {
		return nil
	}
	ds.Start()
	return nil
}
