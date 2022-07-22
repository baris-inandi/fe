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
			&cli.PathFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Install contents from a newline-separated package list file",
			},
		},
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'S')
			cmd.AddFlagsImplicit("--config ~/.config/fe/paru.conf", "removemake", "skipreview", "needed")
			file := c.String("file")
			if file != "" {
				cmd.FormWithSubstitute("- <" + file)
			} else {
				cmd.FormWithArgs()
			}
			cmd.Exec()
			return nil
		},
	}
}
