package cmd

import (
	"fmt"
)

type passwordParam struct {
	name        string `schema.org: "/name"`
	description string `schema.org: "/description"`
	isCommand   bool
}

// NewPasswordParam returns a passwordParam with default values.
func NewPasswordParam() *passwordParam {
	p := &passwordParam{
		name:        "password",
		description: "Generate a random password.``",
		isCommand:   true,
	}
	return p
}

func (p *passwordParam) Name() string {
	return p.name
}

func (p *passwordParam) Description() string {
	return p.description
}

func (p *passwordParam) IsCommand() bool {
	return p.isCommand
}

func (p *passwordParam) String() string {
	return fmt.Sprint(*p)
}

func (p *passwordParam) Set(value string) (err error) {
	return nil
}

func (p *passwordParam) IsBoolFlag() bool { return true }
