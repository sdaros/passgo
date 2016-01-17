package cmd

import (
	"github.com/sdaros/passgo/environment"
)

type (
	// Executable is implemented by all commands in passgo.
	Executable interface {
		Execute(*environment.Env) (CommandResult, error)
		Name() string
	}
	// ExecuteFn holds the Execute() method from an Executable Command.
	ExecuteFn func(*environment.Env) (CommandResult, error)
	// CommandResult returned by a command.
	CommandResult interface{}
)

var (
	passgoCommands = []Executable{
		NewPassword(),
	}
	PassgoCommands = make(map[string]Executable)
)

func init() {
	for _, cmd := range passgoCommands {
		cmdName := cmd.Name()
		PassgoCommands[cmdName] = cmd
	}
}
