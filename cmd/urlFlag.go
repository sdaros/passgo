package cmd

import (
	"fmt"
	"unicode/utf8"
)

type urlFlag struct {
	name      string `schema.org: "/name"`
	usage     string `schema.org: "/description"`
	value     string `schema.org: "/value"`
	isCommand bool
}

// NewUrlFlag returns a new empty UserNameFlag.
func NewUrlFlag() *urlFlag {
	return &urlFlag{
		name:      "url",
		usage:     "Url associated with a secret.",
		isCommand: false,
	}
}

func (ur *urlFlag) Name() string {
	return ur.name
}

func (ur *urlFlag) Usage() string {
	return ur.usage
}

func (ur *urlFlag) IsCommand() bool {
	return ur.isCommand
}

// String is provided to satisfy flag.Value interface.
func (ur *urlFlag) String() string {
	return fmt.Sprint(ur.value)
}

// Set sets the value for the urlFlag and validates it.
func (ur *urlFlag) Set(fromCli string) (err error) {
	if err := ur.Validate(fromCli); err != nil {
		return err
	}
	ur.value = fromCli
	return nil
}

// Validate that flag value is a valid utf8 string.
func (ur *urlFlag) Validate(fromCli string) (err error) {
	if !utf8.ValidString(fromCli) {
		return fmt.Errorf("Expected url to be a valid utf8 string,"+
			" got %v instead", fromCli)
	}
	return nil
}
