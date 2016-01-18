package cmd

import (
	"fmt"
)

type generateCommandFlag struct {
	name      string `schema.org: "/name"`
	usage     string
	isCommand bool
}

// NewGenerateCommandFlag returns a generateCommandFlag with default values.
func NewGenerateCommandFlag() *generateCommandFlag {
	return &generateCommandFlag{
		name:      "generate",
		usage:     "Generate a new sealed secret.",
		isCommand: true,
	}
}

func (g *generateCommandFlag) Name() string {
	return g.name
}

func (g *generateCommandFlag) Usage() string {
	return g.usage
}

func (g *generateCommandFlag) IsCommand() bool {
	return g.isCommand
}

func (g *generateCommandFlag) String() string {
	return fmt.Sprint(*g)
}

func (g *generateCommandFlag) Set(value string) (err error) {
	return nil
}

func (g *generateCommandFlag) IsBoolFlag() bool { return true }
