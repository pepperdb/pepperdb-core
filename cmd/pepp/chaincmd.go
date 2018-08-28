package main

import (
	"fmt"
	"strconv"

	"bytes"
	"encoding/json"

	"github.com/pepperdb/pepperdb-core/core"
	"github.com/urfave/cli"
)

var (
	initCommand = cli.Command{
		Action:    MergeFlags(initGenesis),
		Name:      "init",
		Usage:     "Bootstrap and initialize a new genesis block",
		ArgsUsage: "<genesisPath>",
		Category:  "BLOCKCHAIN COMMANDS",
		Description: `
The init command initializes a new genesis block and definition for the network.`,
	}
	genesisCommand = cli.Command{
		Name:     "genesis",
		Usage:    "the genesis block command",
		Category: "BLOCKCHAIN COMMANDS",
		Description: `
The genesis command for genesis dump or other commands.`,
		Subcommands: []cli.Command{
			{
				Name:   "dump",
				Usage:  "dump the genesis",
				Action: MergeFlags(dumpGenesis),
				Description: `
    pepp account new

Dump the genesis config info.`,
			},
		},
	}

	blockDumpCommand = cli.Command{
		Action:    MergeFlags(dumpblock),
		Name:      "dump",
		Usage:     "Dump the number of newest block before tail block from storage",
		ArgsUsage: "<blocknumber>",
		Category:  "BLOCKCHAIN COMMANDS",
		Description: `
Use "./pepp dump 10" to dump 10 blocks before tail block.`,
	}
)

func initGenesis(ctx *cli.Context) error {
	filePath := ctx.Args().First()
	genesis, err := core.LoadGenesisConf(filePath)
	if err != nil {
		FatalF("load genesis conf faild: %v", err)
	}

	neb, err := makePepp(ctx)
	if err != nil {
		return err
	}

	neb.SetGenesis(genesis)
	neb.Setup()
	return nil
}

func dumpGenesis(ctx *cli.Context) error {
	neb, err := makePepp(ctx)
	if err != nil {
		FatalF("dump genesis conf faild: %v", err)
	}

	neb.Setup()

	genesis, err := core.DumpGenesis(neb.BlockChain())
	if err != nil {
		FatalF("dump genesis conf faild: %v", err)
	}
	genesisJSON, err := json.Marshal(genesis)

	var buf bytes.Buffer
	err = json.Indent(&buf, genesisJSON, "", "    ")
	if err != nil {
		FatalF("dump genesis conf faild: %v", err)
	}
	fmt.Println(buf.String())
	return nil
}

func dumpblock(ctx *cli.Context) error {
	neb, err := makePepp(ctx)
	if err != nil {
		return err
	}

	neb.Setup()

	count, err := strconv.Atoi(ctx.Args().First())
	if err != nil {
		return err
	}
	fmt.Printf("blockchain dump: %s\n", neb.BlockChain().Dump(count))
	return nil
}
