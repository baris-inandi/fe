package command

import (
	"strings"

	"github.com/baris-inandi/fe/utils"
	"github.com/urfave/cli/v2"
)

type Cmd struct {
	operation        rune     // S
	options          []rune   // yu
	flags            []string // -c --quiet
	existingFlagKeys []string // for internal use
}

func makeFlag(f string) string {
	if len(f) < 1 {
		return ""
	}
	if !strings.HasPrefix(f, "-") {
		if len(f) == 1 {
			return "-" + f
		} else {
			return "--" + f
		}
	}
	return f
}

func New(operation rune) Cmd {
	return Cmd{
		operation:        operation,
		options:          []rune{},
		flags:            []string{},
		existingFlagKeys: []string{},
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
	flag = makeFlag(flag)
	if !utils.StringSliceContains(c.flags, flag) {
		c.flags = append(c.flags, flag)
	}
}

func (c *Cmd) AddFlags(flags ...string) {
	for _, flag := range flags {
		c.AddFlag(flag)
	}
}

func (c *Cmd) AddStringFlag(flag string, value string) {
	flag = makeFlag(flag)
	if !utils.StringSliceContains(c.existingFlagKeys, flag) && !utils.StringSliceContains(c.flags, flag) {
		c.existingFlagKeys = append(c.existingFlagKeys, flag)
		c.flags = append(c.flags, flag+" "+value)
	}
}

func (c *Cmd) AddOption(opt rune) {
	if !utils.RuneSliceContains(c.options, opt) {
		c.options = append(c.options, opt)
	}
}

func (c *Cmd) AddOptions(options ...rune) {
	for _, opt := range options {
		c.AddOption(opt)
	}
}
