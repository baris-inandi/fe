package command

import (
	"os"
	"os/exec"
	"strings"

	"github.com/baris-inandi/fe/utils"
)

func (c *Cmd) FormNoArgs() string {
	ctx := c.ctx
	arbitraryFlags := strings.Split(ctx.String("paru"), " ")
	for _, a := range arbitraryFlags {
		c.addFlag(a)
	}
	out := c.bin + " -" + string(c.operation) + string(c.options) + " " + strings.Join(c.flags, " ")
	c.Command = out
	return out
}

func (c *Cmd) FormWithSubstitute(substitution string) string {
	out := c.FormNoArgs() + " " + substitution
	c.Command = out
	return out
}

func (c *Cmd) FormWithArgs() string {
	ctx := c.ctx
	arg := " " + strings.Join(ctx.Args().Slice(), " ")
	out := c.FormNoArgs() + arg
	c.Command = out
	return out
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
	if !ctx.Bool("default") && !ctx.Bool("noparu") {
		for _, flag := range flags {
			c.addFlag(flag)
		}
	}
}

func (c *Cmd) SetBinary(bin string) {
	c.bin = bin
}

func feExec(str string) error {
	cmd := exec.Command("bash", "-c", str)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func StrExec(str string) error {
	return feExec(str)
}

func (c *Cmd) ExecHandleErr() error {
	return feExec(c.Command)
}

func (c *Cmd) Exec() {
	err := c.ExecHandleErr()
	if err != nil {
		os.Exit(1)
	}
}
