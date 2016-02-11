package cmd

import (
	"encoding/json"
	"github.com/sdaros/passgo/app"
)

type (
	// Command supported by passgo.
	Command interface {
		Name() string
		ExecuteFn() func() (*CmdResult, error)
		ApplyCommandFlags(*app.App)
	}
	CmdResult struct {
		Value interface{}
	}
)

var (
	passgoCommands = []Command{
		NewPassword(),
	}
	PassgoCommands = make(map[string]Command)
)

func (c *CmdResult) String() string {
	switch c.Value.(type) {
	case string:
		return c.Value.(string)
	default:
		return ""
	}
}

func (c *CmdResult) Jsonify() (string, error) {
	jsonResult, err := json.MarshalIndent(c, " ", "\t")
	if err != nil {
		return "", nil
	}
	return string(jsonResult), nil

}

func init() {
	for _, cmd := range passgoCommands {
		PassgoCommands[cmd.Name()] = cmd
	}
}
