package fecli

import "github.com/urfave/cli/v2"

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "aur",
			Usage: "Assume targets are from the AUR",
		},
		&cli.BoolFlag{
			Name:  "noaur",
			Usage: "Assume targets are from the official repos",
		},
		&cli.BoolFlag{
			Name:  "review",
			Usage: "Review packages before making a change",
		},
		&cli.StringFlag{
			Name:  "paru",
			Usage: "Pass options to paru",
		},
		&cli.PathFlag{
			Name:  "config",
			Usage: "Overrides fe config file to given path",
		},
		&cli.BoolFlag{
			Name:  "help-paru",
			Usage: "Shows paru help message",
		},
	}
}
