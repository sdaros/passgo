package cmd

import (
	"fmt"
	"unicode/utf8"
)

type userNameFlag struct {
	name      string `schema.org: "/name"`
	usage     string `schema.org: "/description"`
	value     string `schema.org: "/value"`
	isCommand bool
}

// NewUserNameFlag returns a new empty UserNameFlag.
func NewUserNameFlag() *userNameFlag {
	return &userNameFlag{
		name:      "user-name",
		usage:     "Username associated with a secret.",
		isCommand: false,
	}
}

func (un *userNameFlag) Name() string {
	return un.name
}

func (un *userNameFlag) Usage() string {
	return un.usage
}

func (un *userNameFlag) IsCommand() bool {
	return un.isCommand
}

// String is provided to satisfy flag.Value interface.
func (un *userNameFlag) String() string {
	return fmt.Sprint(un.value)
}

// Set sets the value for the userNameFlag and validates it.
func (un *userNameFlag) Set(fromCli string) (err error) {
	if err := un.Validate(fromCli); err != nil {
		return err
	}
	un.value = fromCli
	return nil
}

// Validate that flag value is a valid utf8 string.
func (un *userNameFlag) Validate(fromCli string) (err error) {
	if !utf8.ValidString(fromCli) {
		return fmt.Errorf("Expected username to be a valid utf8 string,"+
			" got %q instead", fromCli)
	}
	return nil
}
