package subcommands

import (
	"fmt"

	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func Search() *cli.Command {
	return &cli.Command{
		Name:    "search",
		Aliases: []string{"se"},
		Usage:   "Searches for packages using given keyword",
		Action: func(c *cli.Context) error {
			cmd := command.New('S')
			cmd.AddOption('s')
			x := cmd.Form(c)
			fmt.Println(x)
			return nil
		},
	}
}
