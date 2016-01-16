package cmd

import (
	"github.com/sdaros/passgo/environment"
	"reflect"
)

type (
	// Executable is implemented by all commands in passgo.
	Executable interface {
		Execute(*environment.Env) (CommandResult, error)
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
		cmdName := reflect.ValueOf(cmd).Elem().FieldByName("name").String()
		PassgoCommands[cmdName] = cmd
	}
}
