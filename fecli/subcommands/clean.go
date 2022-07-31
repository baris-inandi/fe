package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Clean() *cli.Command {
	return &cli.Command{
		Name:    "clean",
		Aliases: []string{"cl"},
		Usage:   "Removes unnecessary packages",
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'R')
			cmd.AddOptions('s')
			cmd.AddFlagsImplicit("skipreview", "cleanafter")
			cmd.FormWithSubstitute("$(paru -Qqtd)")
			cmd.Exec()
			return nil
		},
	}
}
