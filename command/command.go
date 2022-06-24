package command

import (
	"strings"

	"github.com/baris-inandi/fe/utils"
	"github.com/urfave/cli/v2"
)

type Cmd struct {
	operation rune     // S
	options   []rune   // yu
	flags     []string // -c --quiet
}

func New(operation rune) Cmd {
	return Cmd{
		operation: operation,
		options:   []rune{},
		flags:     []string{},
	}
}

func (c *Cmd) Form(ctx *cli.Context) string {
	args := strings.Split(ctx.String("paru"), " ")
	for _, arg := range args {
		c.AddFlag(arg)
	}
	return "paru -" + string(c.operation) + string(c.options) + " " + strings.Join(c.flags, " ")
}

func (c *Cmd) AddFlag(flag string) {
	if !utils.StringContains(c.flags, flag) {
		if !strings.HasPrefix(flag, "-") {
			if len(flag) == 1 {
				flag = "-" + flag
			} else {
				flag = "--" + flag
			}
		}
		c.flags = append(c.flags, flag)
	}
}

func (c *Cmd) AddOption(opt rune) {
	if !utils.RuneContains(c.options, opt) {
		c.options = append(c.options, opt)
	}
}
