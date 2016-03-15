package password

import (
	"fmt"
	"strconv"
)

type NoSymbols struct {
	name  string
	usage string
	value bool
}

// NewNoSymbolsFlag returns a NoSymbols parameter with default values.
func NewNoSymbols() *NoSymbols {
	return &NoSymbols{
		name:  "no-symbols",
		usage: "Do not use special symbols",
		value: false,
	}
}

func (ns *NoSymbols) Name() string {
	return ns.name
}

func (ns *NoSymbols) Usage() string {
	return ns.usage
}

func (ns *NoSymbols) String() string {
	return fmt.Sprint(ns.value)
}

func (ns *NoSymbols) Value() bool { return ns.value }

func (ns *NoSymbols) Set(value string) (err error) {
	noSymb, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	ns.value = noSymb
	return nil
}

func (ns *NoSymbols) Validate() (err error) { return nil }

func (ns *NoSymbols) IsBoolFlag() bool { return true }
