package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Fuzzy() *cli.Command {
	return &cli.Command{
		Name:    "fuzzy",
		Aliases: []string{"fz"},
		Usage:   "Fuzzy-search for packages using paruz",
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'S')
			cmd.AddFlagsImplicit("removemake", "skipreview", "needed")
			cmd.SetBinary("paruz")
			cmd.FormNoArgs()
			cmd.Exec()
			return nil
		},
	}
}
