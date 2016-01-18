package cmd

import (
	"fmt"
	"strconv"
)

type noSymbolsFlag struct {
	name      string `schema.org: "/name"`
	usage     string
	value     bool `schema.org: "/value"`
	isCommand bool
}

// NewNoSymbolsFlag returns a noSymbols parameter with default values.
func NewNoSymbolsFlag() *noSymbolsFlag {
	return &noSymbolsFlag{
		name:      "no-symbols",
		usage:     "Use only alphabetic characters",
		value:     false,
		isCommand: false,
	}
}

func (ns *noSymbolsFlag) Name() string {
	return ns.name
}

func (ns *noSymbolsFlag) Usage() string {
	return ns.usage
}

func (ns *noSymbolsFlag) IsCommand() bool {
	return ns.isCommand
}

func (ns *noSymbolsFlag) String() string {
	return fmt.Sprint(*ns)
}

func (ns *noSymbolsFlag) Set(value string) (err error) {
	noSymb, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	ns.value = noSymb
	return nil
}
