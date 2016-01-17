package cmd

import (
	"fmt"
	"strconv"
)

type passwordLengthFlag struct {
	name      string `schema.org: "/name"`
	usage     string `schema.org: "/description"`
	value     int    `schema.org: "/value"`
	isCommand bool
}

// NewPasswordLengthFlag returns a passwordLength parameter with default values.
func NewPasswordLengthFlag() *passwordLengthFlag {
	pl := &passwordLengthFlag{
		name:      "password-length",
		usage:     "Length of password to be generated.",
		value:     15,
		isCommand: false,
	}
	return pl
}

func (pl *passwordLengthFlag) Name() string {
	return pl.name
}

func (pl *passwordLengthFlag) Usage() string {
	return pl.usage
}

func (pl *passwordLengthFlag) IsCommand() bool {
	return pl.isCommand
}

// String is provided to satisfy flag.Value interface.
func (pl *passwordLengthFlag) String() string {
	return fmt.Sprint(*pl)
}

// Set sets the value for the passwordLengthFlag and validates the range.
func (pl *passwordLengthFlag) Set(value string) (err error) {
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

func (pl *passwordLengthFlag) Validate(length int) (err error) {
	const passwordLengthMin = 1
	const passwordLengthMax = 256
	if length < passwordLengthMin || length > passwordLengthMax {
		err = fmt.Errorf("Password length must be between %v and %v characters",
			passwordLengthMin, passwordLengthMax)
		return err
	}
	return nil
}
