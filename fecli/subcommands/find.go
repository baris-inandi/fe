package subcommands

import (
	"github.com/urfave/cli/v2"
)

func Find() *cli.Command {
	return &cli.Command{
		Name:    "find",
		Aliases: []string{"fd"},
		Usage:   "Fuzzy-search for packages in a tui",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
