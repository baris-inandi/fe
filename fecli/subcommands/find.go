package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Find() *cli.Command {
	return &cli.Command{
		Name:    "find",
		Aliases: []string{"fd"},
		Usage:   "Fuzzy-search for packages using paruz",
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'S')
			cmd.AddFlagsImplicit("--config ~/.config/fe/paru.conf", "removemake", "skipreview", "needed")
			cmd.SetBinary("paruz")
			cmd.FormNoArgs()

			return nil
		},
	}
}
