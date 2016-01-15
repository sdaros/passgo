package cmd

import (
	"github.com/sdaros/passgo/environment"
	"reflect"
)

// Command supported by passgo.
type Command interface {
	Execute(*environment.Env) (CommandResult, error)
}

type CommandResult interface {
	// CommandResult returned by a command.
}

var passgoCommands = []Command{
	NewPassword(),
}
var PassgoCommands = make(map[string]Command)

func init() {
	for _, cmd := range passgoCommands {
		cmdName := reflect.ValueOf(cmd).Elem().FieldByName("name").String()
		PassgoCommands[cmdName] = cmd
	}
}
