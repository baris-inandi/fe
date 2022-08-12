package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Checkupdates() *cli.Command {
	return &cli.Command{
		Name:     "checkupdates",
		Category: CATEGORY_UTILS,
		Aliases:  []string{"cu"},
		Usage:    "Prints a list of pending updates",
		Action: func(c *cli.Context) error {
			return command.StrExec("checkupdates")
		},
	}
}
