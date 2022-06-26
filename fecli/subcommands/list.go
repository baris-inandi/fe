package subcommands

import (
	"fmt"

	"github.com/baris-inandi/fe/command"
	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "Queries packages",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "explicit",
				Aliases: []string{"e"},
				Usage:   "Only list explicitly installed packages (don't include dependencies)",
			},
			&cli.BoolFlag{
				Name:    "deps",
				Aliases: []string{"d"},
				Usage:   "Only list packages installed as dependencies",
			},
			&cli.BoolFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "Outputs package path instead of package version",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Only output package name",
			},
		},
		Action: func(c *cli.Context) error {
			cmd := command.New(c, 'Q')
			if c.Bool("path") {
				cmd.AddOptions('l')
			}
			if c.Bool("quiet") {
				cmd.AddOptions('q')
			}
			if c.Bool("explicit") {
				cmd.AddOptions('e')
			}
			if c.Bool("deps") {
				cmd.AddOptions('d')
			}
			cmd.FormNoArgs()
			fmt.Println(cmd.Command)
			return nil
		},
	}
}
