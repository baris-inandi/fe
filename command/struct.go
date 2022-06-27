package command

import (
	"github.com/urfave/cli/v2"
)

type Cmd struct {
	operation        rune         // S
	options          []rune       // yu
	flags            []string     // -c --quiet
	existingFlagKeys []string     // for internal use
	ctx              *cli.Context // cli context for internal use
	Command          string       // the output command after forming
}

func New(ctx *cli.Context, operation rune) Cmd {
	return Cmd{
		operation:        operation,
		options:          []rune{},
		flags:            []string{},
		existingFlagKeys: []string{},
		ctx:              ctx,
	}
}
