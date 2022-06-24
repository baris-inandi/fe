package flags

import "github.com/urfave/cli/v2"

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "paru",
			Usage: "Pass arbitrary options to paru",
		},
		&cli.PathFlag{
			Name:  "config",
			Usage: "Overrides fe config file to given path",
		},
	}
}
