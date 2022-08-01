package subcommands

import (
	"fmt"
	"os"
	"os/user"

	"github.com/urfave/cli/v2"
)

func Configure() *cli.Command {
	return &cli.Command{
		Name:    "configure",
		Aliases: []string{"cf"},
		Usage:   "Overwrite /etc/paru.conf with recommended configuration",
		Action: func(c *cli.Context) error {
			usr, _ := user.Current()
			if usr.Username != "root" {
				fmt.Println("You cannot perform this operation unless you are root.")
				os.Exit(1)
			}
			conf1 := "/etc/paru.conf"
			if _, err := os.Stat(conf1); err == nil {
				fmt.Println("Removing old paru config file", conf1)
				os.Remove(conf1)
			}
			fmt.Println("Creating symlink: /etc/paru.conf -> /etc/feparu.conf")
			err := os.Symlink("/etc/feparu.conf", "/etc/paru.conf")
			if err != nil {
				fmt.Println("Failed to create symlink: ")
				fmt.Println(err)
				return err
			}
			fmt.Println("Successfully configured fe.")
			return nil
		},
	}
}
