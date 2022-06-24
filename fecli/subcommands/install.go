package subcommands

import (
	"github.com/baris-inandi/fe/actions"
	"github.com/urfave/cli/v2"
)

func Install() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"in"},
		Usage:   "Syncs packages",
		Action: func(c *cli.Context) error {
			return actions.Install(c)
		},
	}
}
