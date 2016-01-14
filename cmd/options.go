package cmd

// options defines all command line flags and validates input
// by implementing the flag.Value interface.

import (
	"fmt"
	"github.com/sdaros/passgo/environment"
	"strconv"
)

type Options []interface {
	// options available for all commands
}

func RegisterOptions(env *environment.Env) {
	var registeredOptions Options
	registeredOptions = []interface{}{
		&passwordLength{
			name:        "password-length",
			description: "Length of password to be generated.",
			value:       15,
		},
		&noSymbols{
			name:        "no-symbols",
			description: "Use only alphabetic characters.",
			value:       false,
		},
	}
	env.Register("options", registeredOptions)
}

// passwordLength option.
type passwordLength struct {
	name        string
	description string
	value       int
}

// NewPasswordLength returns a passwordLength option with default values.
func NewPasswordLength() *passwordLength {
	pl := &passwordLength{
		name:        "password-length",
		description: "Length of password to be generated.",
		value:       15,
	}
	return pl
}

// String is provided to satisfy flag.Value interface.
func (pl *passwordLength) String() string {
	return fmt.Sprint(*pl)
}

// Set sets the value for the passwordLength option and validates the range.
func (pl *passwordLength) Set(value string) (err error) {
	length, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := pl.Validate(length); err != nil {
		return err
	}
	pl.value = length
	return nil
}

func (pl *passwordLength) Validate(length int) (err error) {
	const passwordLengthMin = 1
	const passwordLengthMax = 256
	if length < passwordLengthMin || length > passwordLengthMax {
		err = fmt.Errorf("Password length must be between %v and %v characters",
			passwordLengthMin, passwordLengthMax)
		return err
	}
	return nil
}

// noSymbols option.
type noSymbols struct {
	name        string
	description string
	value       bool
}

// NewPasswordLength returns a passwordLength option with default values.
func NewNoSymbols() *noSymbols {
	ns := &noSymbols{
		name:        "no-symbols",
		description: "Use only alphabetic characters",
		value:       false,
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
