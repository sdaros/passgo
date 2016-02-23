package cmd

import (
	"github.com/sdaros/passgo/app"
)

type (
	// Command supported by passgo.
	Command interface {
		Name() string
		ExecuteFn() func() (CmdResult, error)
		ApplyCommandFlagsFrom(*app.App) error
	}
	CmdResult interface{}
)

var (
	passgoCommands = []Command{
		NewPassword(),
		NewGenerate(),
	}
	PassgoCommands = make(map[string]Command)
)

func init() {
	for _, cmd := range passgoCommands {
		PassgoCommands[cmd.Name()] = cmd
	}
}
