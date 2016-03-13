package cmd

import (
	"github.com/sdaros/passgo/app"
)

type (
	// Command supported by passgo.
	Command interface {
		ApplyCommandFlagsFrom(*app.App) error
		ExecuteFn() func() (CmdResult, error)
		Name() string
	}
	Param interface {
		Name() string
		Usage() string
		Validate(interface{}) error
	}
	CmdResult interface{}
)

var (
	passgoCommands = []Command{
		NewPassword(),
		NewGenerate(),
		NewStamp(),
		NewScrypt(),
	}
	PassgoCommands = make(map[string]Command)
)

func init() {
	for _, cmd := range passgoCommands {
		PassgoCommands[cmd.Name()] = cmd
	}
}
