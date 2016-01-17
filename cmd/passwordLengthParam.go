package cmd

import (
	"fmt"
	"strconv"
)

type passwordLengthParam struct {
	name        string `schema.org: "/name"`
	description string `schema.org: "/description"`
	value       int    `schema.org: "/value"`
	isCommand   bool
}

// NewPasswordLengthParam returns a passwordLength parameter with default values.
func NewPasswordLengthParam() *passwordLengthParam {
	pl := &passwordLengthParam{
		name:        "password-length",
		description: "Length of password to be generated.",
		value:       15,
		isCommand:   false,
	}
	return pl
}

func (pl *passwordLengthParam) Name() string {
	return pl.name
}

func (pl *passwordLengthParam) Description() string {
	return pl.description
}

func (pl *passwordLengthParam) IsCommand() bool {
	return pl.isCommand
}

// String is provided to satisfy flag.Value interface.
func (pl *passwordLengthParam) String() string {
	return fmt.Sprint(*pl)
}

// Set sets the value for the passwordLengthParam and validates the range.
func (pl *passwordLengthParam) Set(value string) (err error) {
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

func (pl *passwordLengthParam) Validate(length int) (err error) {
	const passwordLengthMin = 1
	const passwordLengthMax = 256
	if length < passwordLengthMin || length > passwordLengthMax {
		err = fmt.Errorf("Password length must be between %v and %v characters",
			passwordLengthMin, passwordLengthMax)
		return err
	}
	return nil
}
