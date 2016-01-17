package cmd

import (
	"fmt"
	"strconv"
)

type noSymbolsParam struct {
	name        string `schema.org: "/name"`
	description string `schema.org: "/description"`
	value       bool   `schema.org: "/value"`
	isCommand   bool
}

// NewNoSymbolsParam returns a noSymbols parameter with default values.
func NewNoSymbolsParam() *noSymbolsParam {
	ns := &noSymbolsParam{
		name:        "no-symbols",
		description: "Use only alphabetic characters",
		value:       false,
		isCommand:   false,
	}
	return ns
}

func (ns *noSymbolsParam) Name() string {
	return ns.name
}

func (ns *noSymbolsParam) Description() string {
	return ns.name
}

func (ns *noSymbolsParam) IsCommand() bool {
	return ns.isCommand
}

func (ns *noSymbolsParam) String() string {
	return fmt.Sprint(*ns)
}

func (ns *noSymbolsParam) Set(value string) (err error) {
	noSymb, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	ns.value = noSymb
	return nil
}
