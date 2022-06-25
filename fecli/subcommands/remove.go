package subcommands

import (
	"fmt"

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
			cmd := command.New('R')
			if c.Bool("query") {
				// TODO: implement the following script:
				/*
				   set targets (paru -Qqs $argv)
				   if test -n "$targets"
				       paru -Rs $targets --skipreview --cleanafter --removemake
				   else
				       echo "Query returned no results."
				       return 1
				   end
				*/
			}
			cmd.AddFlags("skipreview", "noconfirm", "cleanafter")
			if !c.Bool("keep-deps") {
				cmd.AddOption('s')
			}
			a := cmd.Form(c)
			fmt.Println(a)
			return nil
		},
	}
}
