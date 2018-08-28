package main

import (
	"fmt"

	"github.com/pepperdb/pepperdb-core/neblet"
	"github.com/urfave/cli"
)

var (
	configCommand = cli.Command{
		Name:     "config",
		Usage:    "Manage config",
		Category: "CONFIG COMMANDS",
		Description: `
Manage neblas config, generate a default config file.`,

		Subcommands: []cli.Command{
			{
				Name:      "new",
				Usage:     "Generate a default config file",
				Action:    MergeFlags(createDefaultConfig),
				ArgsUsage: "<filename>",
				Description: `
Generate a a default config file.`,
			},
		},
	}
)

// accountCreate creates a new account into the keystore
func createDefaultConfig(ctx *cli.Context) error {
	fileName := ctx.Args().First()
	if len(fileName) == 0 {
		fmt.Println("please give a config file arg!!!")
		return nil
	}
	neblet.CreateDefaultConfigFile(fileName)
	fmt.Printf("create default config %s\n", fileName)
	return nil
}
