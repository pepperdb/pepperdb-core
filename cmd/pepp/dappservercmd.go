package main

import (
	"github.com/pepperdb/pepperdb-core/dappserver"
	"github.com/urfave/cli"
)

var (
	// DAppServerConfigFlag dappserver config file path
	DAppServerConfigFlag = cli.StringFlag{
		Name:        "config, c",
		Usage:       "load dappserver configuration from `FILE`",
		Value:       "conf/example/dappserver.conf",
		Destination: &config,
	}

	dappServerCommand = cli.Command{
		Action:    MergeFlags(startDAppServer),
		Name:      "dappserver",
		Usage:     "Start dapp server",
		ArgsUsage: "<dapp server config>",
		Category:  "DAPPSERVER COMMANDS",
		Description: `
The dappserver command start a dapp server.`,
		Flags: []cli.Flag{
			DAppServerConfigFlag,
		},
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
