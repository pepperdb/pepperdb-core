package main

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"
	"time"

	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/core"
	"github.com/pepperdb/pepperdb-core/neblet"
	"github.com/urfave/cli"
)

var (
	version   string
	commit    string
	branch    string
	compileAt string
	config    string
)

func main() {
	app := cli.NewApp()
	app.Action = pepp
	app.Name = "PepperDB Core"
	app.Version = fmt.Sprintf("%s, branch %s, commit %s", version, branch, commit)
	timestamp, _ := strconv.ParseInt(compileAt, 10, 64)
	app.Compiled = time.Unix(timestamp, 0)
	app.Usage = "The PepperDB command line tools"
	app.Copyright = "Copyright 2018-2019 The PepperDB team"

	app.Flags = append(app.Flags, ConfigFlag)
	app.Flags = append(app.Flags, NetworkFlags...)
	app.Flags = append(app.Flags, ChainFlags...)
	app.Flags = append(app.Flags, RPCFlags...)
	app.Flags = append(app.Flags, AppFlags...)
	app.Flags = append(app.Flags, StatsFlags...)

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Commands = []cli.Command{
		initCommand,
		genesisCommand,
		accountCommand,
		consoleCommand,
		networkCommand,
		versionCommand,
		licenseCommand,
		configCommand,
		blockDumpCommand,
		dappServerCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}

func pepp(ctx *cli.Context) error {
	n, err := makePepp(ctx)
	if err != nil {
		return err
	}

	logging.Init(n.Config().App.LogFile, n.Config().App.LogLevel, n.Config().App.LogAge)

	core.SetCompatibilityOptions(n.Config().Chain.ChainId)

	// enable crash report if open the switch and configure the url
	if n.Config().App.EnableCrashReport && len(n.Config().App.CrashReportUrl) > 0 {
		InitCrashReporter(n.Config().App)
	}

	select {
	case <-runPepp(ctx, n):
		return nil
	}
}

func runPepp(ctx *cli.Context, n *neblet.Neblet) chan bool {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// start net pprof if config.App.Pprof.HttpListen configured
	err := n.StartPprof(n.Config().App.Pprof.HttpListen)
	if err != nil {
		FatalF("start pprof failed:%s", err)
	}

	n.Setup()
	n.Start()

	quitCh := make(chan bool, 1)

	go func() {
		<-c

		n.Stop()

		quitCh <- true
		return
	}()

	return quitCh
}

func makePepp(ctx *cli.Context) (*neblet.Neblet, error) {
	conf := neblet.LoadConfig(config)
	conf.App.Version = version

	// load config from cli args
	networkConfig(ctx, conf.Network)
	chainConfig(ctx, conf.Chain)
	rpcConfig(ctx, conf.Rpc)
	appConfig(ctx, conf.App)
	statsConfig(ctx, conf.Stats)

	n, err := neblet.New(conf)
	if err != nil {
		return nil, err
	}
	return n, nil
}

// FatalF fatal format err
func FatalF(format string, args ...interface{}) {
	err := fmt.Sprintf(format, args...)
	fmt.Println(err)
	os.Exit(1)
}
