package subcommands

import (
	"github.com/urfave/cli/v2"
)

func Remove() *cli.Command {
	return &cli.Command{
		Name:    "remove",
		Aliases: []string{"rm"},
		Usage:   "Removes packages",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "keep-deps",
				Aliases: []string{"k"},
				Usage:   "Don't uninstall dependencies",
			},
			&cli.BoolFlag{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Remove packages from query output",
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
