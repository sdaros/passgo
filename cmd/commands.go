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
		&Password{
			name:           "password",
			description:    "Length of password to be generated.",
			noSymbols:      new(noSymbols),
			passwordLength: new(passwordLength),
			Env:            env,
		},
	}
	env.Register("commands", registeredCommands)
}
