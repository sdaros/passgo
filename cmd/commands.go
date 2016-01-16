package cmd

import (
	"github.com/sdaros/passgo/environment"
	"reflect"
)

// Executable is implemented by all commands in passgo.
type Executable interface {
	Execute(*environment.Env) (CommandResult, error)
}

type CommandResult interface {
	// CommandResult returned by a command.
}

var passgoCommands = []Executable{
	NewPassword(),
}
var PassgoCommands = make(map[string]Executable)

func init() {
	for _, cmd := range passgoCommands {
		cmdName := reflect.ValueOf(cmd).Elem().FieldByName("name").String()
		PassgoCommands[cmdName] = cmd
	}
}
