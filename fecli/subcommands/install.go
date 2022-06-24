package subcommands

import (
	"github.com/baris-inandi/fe/command"
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
			cmd := command.New('S')
			cmd.AddOption('y')
			cmd.AddOption('u')
			cmd.AddFlag("help")
			cmd.Form(c)
			return nil
		},
	}
}
