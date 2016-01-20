package cmd

import (
	"errors"
	"github.com/sdaros/passgo/environment"
)

var (
	ErrGenerate = errors.New("cmd: Error while trying" +
		"to generate a new secret")
)

// Generate creates a new secret by taking the provided Url and Username and
// appending a randomly generated password.
type Generate struct {
	name           string `schema.org: "/name"`
	passwordLength *passwordLengthFlag
	noSymbols      *noSymbolsFlag
	*environment.Env
}

// NewGenerate returns a secret with default values
func NewGenerate() *Generate {
	return &Generate{
		name:           "generate",
		noSymbols:      NewNoSymbolsFlag(),
		passwordLength: NewPasswordLengthFlag(),
		Env:            environment.Null(),
	}
}
