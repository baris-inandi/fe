package fecli

import (
	"github.com/baris-inandi/fe/fecli/subcommands"
	"github.com/urfave/cli/v2"
)

var Cli = &cli.App{
	Name:    "fe",
	Usage:   "AUR helper with a familiar subcommand system",
	Suggest: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "paru",
			Usage: "Pass arbitrary options to paru",
		},
	},
	Commands: []*cli.Command{
		subcommands.Install(),
		subcommands.Search(),
		subcommands.Update(),
		subcommands.Clean(),
		subcommands.List(),
		subcommands.Remove(),
	},
}

// TODO: use --config to apply the preferred config file
