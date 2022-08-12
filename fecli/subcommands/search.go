package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Search() *cli.Command {
	return &cli.Command{
		Name:     "search",
		Category: CATEGORY_MAIN,
		Aliases:  []string{"se"},
		Usage:    "Searches for packages using given keyword",
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'S')
			cmd.AddOptions('s')
			cmd.FormWithArgs()
			cmd.Exec()
			return nil
		},
	}
}
