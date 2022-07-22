package command

import (
	"github.com/urfave/cli/v2"
)

type Cmd struct {
	bin              string       // paru, paruz, pacman etc.
	operation        rune         // S
	options          []rune       // yu
	flags            []string     // -c --quiet
	existingFlagKeys []string     // for internal use
	ctx              *cli.Context // cli context for internal use
	Command          string       // the output command after forming
}

func New(ctx *cli.Context, operation rune) Cmd {
	bin := "paru"
	if ctx.Bool("noparu") {
		bin = "sudo pacman"
	}
	return Cmd{
		bin:              bin,
		operation:        operation,
		options:          []rune{},
		flags:            []string{},
		existingFlagKeys: []string{},
		ctx:              ctx,
	}
}
