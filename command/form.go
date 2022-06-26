package command

import "strings"

func (c *Cmd) FormNoArgs() string {
	ctx := c.ctx
	arbitraryFlags := strings.Split(ctx.String("paru"), " ")
	for _, a := range arbitraryFlags {
		c.addFlag(a)
	}
	out := "paru -" + string(c.operation) + string(c.options) + " " + strings.Join(c.flags, " ")
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
