package command

import (
	"strings"

	"github.com/baris-inandi/fe/utils"
)

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

func (c *Cmd) addFlag(flag string) {
	flag = makeFlag(flag)
	if !utils.StringSliceContains(c.flags, flag) {
		c.flags = append(c.flags, flag)
	}
}

func (c *Cmd) AddFlags(flags ...string) {
	for _, flag := range flags {
		c.addFlag(flag)
	}
}

func (c *Cmd) AddStringFlag(flag string, value string) {
	flag = makeFlag(flag)
	if !utils.StringSliceContains(c.existingFlagKeys, flag) && !utils.StringSliceContains(c.flags, flag) {
		c.existingFlagKeys = append(c.existingFlagKeys, flag)
		c.flags = append(c.flags, flag+" "+value)
	}
}

func (c *Cmd) AddOptions(options ...rune) {
	for _, opt := range options {
		if !utils.RuneSliceContains(c.options, opt) {
			c.options = append(c.options, opt)
		}
	}
}

func (c *Cmd) AddFlagsImplicit(flags ...string) {
	ctx := c.ctx
	if !ctx.Bool("default") {
		for _, flag := range flags {
			c.addFlag(flag)
		}
	}
}