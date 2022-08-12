package subcommands

import (
	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Depends() *cli.Command {
	return &cli.Command{
		Name:     "depends",
		Category: CATEGORY_UTILS,
		Aliases:  []string{"de"},
		Usage:    "Prints a dependency tree for a package",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "reverse",
				Aliases: []string{"r"},
				Usage:   "List packages that depend on specified package",
			},
		}, Action: func(c *cli.Context) error {
			if c.Bool("reverse") {
				return command.StrExec("pactree -rc " + c.Args().Slice()[0])
			}
			return command.StrExec("pactree -c " + c.Args().Slice()[0])
		},
	}
}
