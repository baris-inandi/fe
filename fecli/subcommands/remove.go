package subcommands

import (
	"fmt"
	"strings"

	"github.com/baris-inandi/fe/command"
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
			cmd := command.New(c, 'R')
			cmd.AddFlagsImplicit("skipreview", "cleanafter")
			if !c.Bool("keep-deps") {
				cmd.AddOptions('s')
			}
			if c.Bool("query") {
				argsSlice := c.Args().Slice()
				args := strings.Join(argsSlice, " ")
				if len(argsSlice) > 0 {
					cmd.FormWithSubstitute("$(paru -Qqs " + args + ")")
					cmd.Exec()
				}
				// TODO: implement the following script:
				/*
				   set targets (paru -Qqs $argv)
				   if test -n "$targets"
				       paru -Rs --skipreview --cleanafter --removemake $(paru -Qqs $argv)
				   else
				       echo "Query returned no results."
				       return 1
				   end
				*/
			}
			a := cmd.FormWithArgs()
			fmt.Println(a)
			return nil
		},
	}
}
