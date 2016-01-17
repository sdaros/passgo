package cmd

import (
	"fmt"
)

type passwordCommandFlag struct {
	name      string `schema.org: "/name"`
	usage     string
	isCommand bool
}

// NewPasswordCommandFlag returns a passwordCommandFlag with default values.
func NewPasswordCommandFlag() *passwordCommandFlag {
	p := &passwordCommandFlag{
		name:      "password",
		usage:     "Generate a random password.``",
		isCommand: true,
	}
	return p
}

func (p *passwordCommandFlag) Name() string {
	return p.name
}

func (p *passwordCommandFlag) Usage() string {
	return p.usage
}

func (p *passwordCommandFlag) IsCommand() bool {
	return p.isCommand
}

func (p *passwordCommandFlag) String() string {
	return fmt.Sprint(*p)
}

func (p *passwordCommandFlag) Set(value string) (err error) {
	return nil
}

func (p *passwordCommandFlag) IsBoolFlag() bool { return true }
