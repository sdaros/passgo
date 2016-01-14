package cmd

import (
	"github.com/sdaros/passgo/environment"
)

type Commands []interface {
	// tracks all available commands.
}

func RegisterCommands(env *environment.Env) {
	var registeredCommands Commands
	registeredCommands = []interface{}{
		NewPassword(env),
	}
	env.Register("commands", registeredCommands)
}
