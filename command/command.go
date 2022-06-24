package command

import (
	"strings"

	"github.com/baris-inandi/fe/utils"
)

type Cmd struct {
	Operation string   // S
	Options   []rune   // yu
	Flags     []string // -c --quiet
}

func New(operation string) Cmd {
	return Cmd{
		Operation: operation,
		Options:   []rune{},
		Flags:     []string{},
	}
}

func (c *Cmd) Form() string {
	return "paru -" + c.Operation + string(c.Options) + " " + strings.Join(c.Flags, " ")
}

func (c *Cmd) AddFlag(flag string) {
	if !utils.StringContains(c.Flags, flag) {
		c.Flags = append(c.Flags, flag)
	}
}

func (c *Cmd) AddOption(opt rune) {
	if !utils.RuneContains(c.Options, opt) {
		c.Options = append(c.Options, opt)
	}
}
