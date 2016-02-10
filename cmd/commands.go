package cmd

import (
	"github.com/sdaros/passgo/app"
)

type (
	// Command supported by passgo.
	Command interface {
		Name() string
		ExecuteFn() func() (string, error)
		ApplyCommandFlags(*app.App)
	}
)

var (
	passgoCommands = []Command{
		NewPassword(),
	}
	PassgoCommands = make(map[string]Command)
)

func init() {
	for _, cmd := range passgoCommands {
		PassgoCommands[cmd.Name()] = cmd
	}
}
