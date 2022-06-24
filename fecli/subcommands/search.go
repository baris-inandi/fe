package subcommands

import (
	"github.com/baris-inandi/fe/actions"
	"github.com/urfave/cli/v2"
)

func Search() *cli.Command {
	return &cli.Command{
		Name:    "search",
		Aliases: []string{"sr"},
		Usage:   "Searches for packages using given keyword",
		Action: func(c *cli.Context) error {
			return actions.Search(c)
		},
	}
}
