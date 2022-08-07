package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Depends() *cli.Command {
	return &cli.Command{
		Name:    "depends",
		Aliases: []string{"de"},
		Usage:   "Prints a dependency tree for a package",
		Action: func(c *cli.Context) error {
			return command.StrExec("pactree -c " + c.Args().Slice()[0])
		},
	}
}
