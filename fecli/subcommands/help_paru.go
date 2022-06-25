package subcommands

import (
	"github.com/urfave/cli/v2"
)

func HelpParu() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "Prints paru help message for usage with the 'paru' flag",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
