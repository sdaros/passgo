package cli

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
		name:        "password",
		description: "generate a random password.",
		isCommand:   true,
	}
	return g
}

func (g *generateParam) String() string {
	return fmt.Sprint(*g)
}

func (g *generateParam) Set(value string) (err error) {
	return nil
}
