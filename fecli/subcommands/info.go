package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Info() *cli.Command {
	return &cli.Command{
		Name:    "info",
		Aliases: []string{"if"},
		Usage:   "Print detailed information about a package",
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'S')
			cmd.AddOptions('i')
			cmd.FormWithArgs()
			cmd.Exec()
			return nil
		},
	}
}
