package cmd

// options defines all command line flags and validates input
// by implementing the flag.Value interface.

import (
	"fmt"
	"github.com/sdaros/passgo/environment"
	"strconv"
)

// options available for all commands
type options []interface{}

func registerOptions(env *environment.Env) {
	var registeredOptions options
	registeredOptions = []interface{}{
		&passwordLength{
			"password-length", "Length of password to be generated.", 15,
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

// String is provided to satisfy flag.Value interface.
func (pl *passwordLength) String() string {
	return fmt.Sprint(*pl)
}

// Set sets the value for the passwordLength option and validates the range.
func (pl *passwordLength) Set(value string) (err error) {
	const passwordLengthMin = 1
	const passwordLengthMax = 256
	length, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if length < passwordLengthMin || length > passwordLengthMax {
		err = fmt.Errorf("Password length must be between %v and %v characters",
			passwordLengthMin, passwordLengthMax)
		return err
	}
	pl.value = length
	return nil
}

// noSymbols option.
type noSymbols bool

func (ns *noSymbols) String() string {
	return fmt.Sprint(*ns)
}

func (ns *noSymbols) Set(value string) (err error) {
	noSymb, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	*ns = noSymbols(noSymb)
	return nil
}
