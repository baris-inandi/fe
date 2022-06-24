package subcommands

import (
	"github.com/baris-inandi/fe/actions"
	"github.com/urfave/cli/v2"
)

func Clean() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"cl"},
		Usage:   "Removes unnecessary packages",
		Action: func(c *cli.Context) error {
			return actions.Clean(c)
		},
	}
}
