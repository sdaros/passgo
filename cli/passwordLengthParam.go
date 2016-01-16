package cli

import (
	"fmt"
	"strconv"
)

// passwordLength option.
type passwordLength struct {
	name        string `schema.org: "/name"`
	description string `schema.org: "/description"`
	value       int    `schema.org: "/value"`
	isCommand   bool
}

// NewPasswordLength returns a passwordLength option with default values.
func NewPasswordLength() *passwordLength {
	pl := &passwordLength{
		name:        "password-length",
		description: "Length of password to be generated.",
		value:       15,
		isCommand:   false,
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
