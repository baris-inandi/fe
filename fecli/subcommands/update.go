package subcommands

import (
	"github.com/baris-inandi/fe/actions"
	"github.com/urfave/cli/v2"
)

func Update() *cli.Command {
	return &cli.Command{
		Name:    "update",
		Aliases: []string{"up"},
		Usage:   "Updates all packages",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "ignore",
				Aliases: []string{"i"},
				Usage:   "Don't update given packages",
			},
		},
		Action: func(c *cli.Context) error {
			return actions.Update(c)
		},
	}
}
