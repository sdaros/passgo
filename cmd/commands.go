package cmd

import (
	"encoding/json"
	"github.com/sdaros/passgo/environment"
)

type (
	// Executable is implemented by all commands in passgo.
	Executable interface {
		Execute(*environment.Env) (*CommandResult, error)
		Name() string
	}
	// ExecuteFunc holds the Execute() method from an Executable Command.
	ExecuteFunc func(*environment.Env) (*CommandResult, error)
	// CommandResult returned by a command.
	CommandResult struct {
		Value interface{} `json:"value"`
	}
)

var (
	passgoCommands = []Executable{
		NewPassword(),
	}
	PassgoCommands = make(map[string]Executable)
)

func (c *CommandResult) String() (string, error) {
	jsonResult, err := json.MarshalIndent(c, " ", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonResult), nil
}

func init() {
	for _, cmd := range passgoCommands {
		cmdName := cmd.Name()
		PassgoCommands[cmdName] = cmd
	}
}
