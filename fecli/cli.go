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
			Name:    "paru",
			Aliases: []string{"p"},
			Usage:   "Pass arbitrary options to paru",
		},
		&cli.BoolFlag{
			Name:    "default",
			Aliases: []string{"d"},
			Usage:   "Don't implicitly pass flags to paru",
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
