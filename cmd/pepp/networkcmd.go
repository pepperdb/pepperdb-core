package main

import (
	"fmt"

	"github.com/pepperdb/pepperdb-core/network/net"
	"github.com/pepperdb/pepperdb-core/common/util"
	"github.com/urfave/cli"
)

var (
	networkCommand = cli.Command{
		Name:     "network",
		Usage:    "Manage network",
		Category: "NETWORK COMMANDS",
		Description: `
Manage neblas network, generate a private key for node.`,

		Subcommands: []cli.Command{
			{
				Name:      "ssh-keygen",
				Usage:     "Generate a private key for network node",
				Action:    generatePrivateKey,
				ArgsUsage: "<path>",
				Description: `

Generate a private key for network node.

If the private key of a network node is exist, the nodeID will not change.

Make sure that the seed node should have a private key.`,
			},
		},
	}
)

// accountCreate creates a new account into the keystore
func generatePrivateKey(ctx *cli.Context) error {
	key, err := net.GenerateEd25519Key()
	if err != nil {
		return err
	}

	str, _ := net.MarshalNetworkKey(key)
	fmt.Printf("private.key: %s\n", key)

	path := ctx.Args().First()
	if len(path) == 0 {
		path = net.DefaultPrivateKeyPath
	}

	return util.FileWrite(path, []byte(str), false)
}
