package cli

import (
	"fmt"
	"strconv"
)

// noSymbols option.
type noSymbols struct {
	name        string `schema.org: "/name"`
	description string `schema.org: "/description"`
	value       bool   `schema.org: "/value"`
	isCommand   bool
}

// NewPasswordLength returns a passwordLength option with default values.
func NewNoSymbols() *noSymbols {
	ns := &noSymbols{
		name:        "no-symbols",
		description: "Use only alphabetic characters",
		value:       false,
		isCommand:   false,
	}
	return ns
}

func (ns *noSymbols) String() string {
	return fmt.Sprint(*ns)
}

func (ns *noSymbols) Set(value string) (err error) {
	noSymb, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	ns.value = noSymb
	return nil
}
