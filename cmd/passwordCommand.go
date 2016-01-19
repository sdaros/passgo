package cmd

import (
	"errors"
	"github.com/sdaros/passgo/environment"
)

var (
	// Define a range of characters that can be used when generating random passwords.
	runesNoSymbols = []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B',
		'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
		'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
		'y', 'z'}
	runesWithSymbols = []rune{
		'!', '#', '$', '%', '&', '(', ')', '*', '+', ',', '-', '.',
		'/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':',
		';', '=', '>', '?', '@', 'A', 'B', 'C', 'D', 'E', 'F', 'G',
		'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
		'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', ']', '^', '_', '`',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
		'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
		'y', 'z', '{', '|', '}', '~'}
	ErrPassword = errors.New("cmd: Error while trying" +
		" to generate password")
)

// Password returns a password to the caller
// based on the parameters provided to it.
type Password struct {
	name           string `schema.org: "/name"`
	noSymbols      *noSymbolsFlag
	passwordLength *passwordLengthFlag
	*environment.Env
}

// NewPassword returns a password command with default values
func NewPassword() *Password {
	return &Password{
		name:           "password",
		noSymbols:      NewNoSymbolsFlag(),
		passwordLength: NewPasswordLengthFlag(),
		Env:            environment.Null(),
	}
}

// Execute validates command options then returns a password
// composed of random elements chosen from a rune pool
func (p *Password) Execute(env *environment.Env) (*CommandResult, error) {
	p.Env = env
	p.applyCommandFlags(env)
	if err := p.validate(); err != nil {
		return new(CommandResult), err
	}
	if p.noSymbols.value {
		result, err := p.composePassword(runesNoSymbols)
		if err != nil {
			return new(CommandResult), err
		}
		return &CommandResult{result}, nil
	}
	result, err := p.composePassword(runesWithSymbols)
	if err != nil {
		return new(CommandResult), err
	}
	return &CommandResult{result}, nil
}

// composePassword of passwordLength by selecting random elements
// from an ASCII subset (runePool).
func (p *Password) composePassword(runePool []rune) (string, error) {
	var password []rune
	for i := 0; i < p.passwordLength.value; i++ {
		runeAtIndex, err := p.randomIndexFromRunePool(runePool)
		if err != nil {
			return "", ErrPassword
		}
		password = append(password, runePool[runeAtIndex])
	}
	return string(password), nil
}

// randomIndexFromRunePool returns index from Rune Pool using the
// Int() method from the Entropy implementation defined in the environment
func (p *Password) randomIndexFromRunePool(runePool []rune) (int64, error) {
	return p.Int(len(runePool) - 1)
}

func (p *Password) applyCommandFlags(env *environment.Env) {
	p.passwordLength = env.Lookup("password-length").(*passwordLengthFlag)
	p.noSymbols = env.Lookup("no-symbols").(*noSymbolsFlag)
}

func (p *Password) validate() (err error) {
	length := p.passwordLength.value
	if err := p.passwordLength.Validate(length); err != nil {
		return err
	}
	return nil
}

func (p *Password) Name() string {
	return p.name
}
