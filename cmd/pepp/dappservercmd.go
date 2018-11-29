package main

import (
	"github.com/pepperdb/pepperdb-core/dappserver"
	"github.com/urfave/cli"
)

var (
	dappServerCommand = cli.Command{
		Action:    MergeFlags(startDAppServer),
		Name:      "dappserver",
		Usage:     "Start dapp server",
		ArgsUsage: "<dapp server config>",
		Category:  "DAPPSERVER COMMANDS",
		Description: `
The dappserver command start a dapp server.`,
	}
)

func startDAppServer(ctx *cli.Context) error {
	conf := dappserver.LoadConfig(config)
	ds, err := dappserver.NewServer(conf)
	if err != nil {
		return err
	}
	ds.Start()
	return nil
}
