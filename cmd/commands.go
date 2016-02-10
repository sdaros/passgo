package cmd

import (
	"encoding/json"
)

type (
	// Command supported by passgo.
	Command struct {
		Executable
		Result interface{}
	}
	// Executable is a command in passgo that can be executed.
	Executable interface {
		Execute() error
		ApplyCommandFlags()
		Name() string
	}
)

var (
	passgoCommands = []Executable{
		NewPassword(),
	}
	PassgoCommands = make(map[string]Executable)
)

func (c *Command) String() (string, error) {
	jsonResult, err := json.MarshalIndent(c, " ", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonResult), nil
}

func init() {
	for _, cmd := range passgoCommands {
		PassgoCommands[cmd.Name()] = cmd
	}
}
