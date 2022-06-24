package subcommands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func Install() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Aliases: []string{"in"},
		Usage:   "Syncs packages",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Select and install packages from query output",
			},
		},
		Action: func(c *cli.Context) error {
			// paru -S [args]
			fmt.Println(c.String("paru"))
			return nil
		},
	}
}
