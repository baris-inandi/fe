package subcommands

import (
	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "Lists installed packages",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
