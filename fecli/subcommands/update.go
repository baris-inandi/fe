package subcommands

import (
	"strings"

	"github.com/baris-inandi/fe/command"

	"github.com/urfave/cli/v2"
)

func Update() *cli.Command {
	return &cli.Command{
		Name:     "update",
		Category: CATEGORY_MAIN,
		Aliases:  []string{"up"},
		Usage:    "Updates all packages",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "ignore",
				Aliases: []string{"i"},
				Usage:   "Don't update given packages",
			},
		},
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'S')
			cmd.AddOptions('y', 'u')
			cmd.AddFlagsImplicit("skipreview", "needed")
			ignoreSlice := c.StringSlice("ignore")
			if len(ignoreSlice) > 0 {
				cmd.AddStringFlag("ignore", strings.Join(ignoreSlice, ","))
			}
			cmd.FormNoArgs()
			cmd.Exec()
			return nil
		},
	}
}
