package cmd

import (
	"fmt"
)

type generateParam struct {
	name        string `schema.org: "/name"`
	description string `schema.org: "/description"`
	isCommand   bool
}

// NewGenerateParam returns a generateParam with default values.
func NewGenerateParam() *generateParam {
	g := &generateParam{
		name:        "generate",
		description: "Generate a new sealed secret.",
		isCommand:   true,
	}
	return g
}

func (g *generateParam) Name() string {
	return g.name
}

func (g *generateParam) Description() string {
	return g.description
}

func (g *generateParam) IsCommand() bool {
	return g.isCommand
}

func (g *generateParam) String() string {
	return fmt.Sprint(*g)
}

func (g *generateParam) Set(value string) (err error) {
	return nil
}
