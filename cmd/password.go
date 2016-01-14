package cmd

import (
	"errors"
	"github.com/sdaros/passgo/environment"
	_ "reflect"
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
		'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',',
		'-', '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8',
		'9', ':', ';', '=', '>', '?', '@', 'A', 'B', 'C', 'D', 'E',
		'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q',
		'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '[', '\\', ']',
		'^', '_', '`', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
		'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
		'v', 'w', 'x', 'y', 'z', '{', '|', '}', '~'}
	ErrPassword = errors.New("cryptoRand: Error while trying" +
		" to generate password")
)

// Password returns a password to the caller
// based on the parameters provided to it.
type Password struct {
	name           string
	description    string
	noSymbols      *noSymbols
	passwordLength *passwordLength
	*environment.Env
}

// Execute returns a password composed of random elements
// chosen from a rune pool
func (p *Password) Execute() (password []rune, err error) {
	if p.noSymbols.value {
		return p.composePassword(runesNoSymbols)
	}
	return p.composePassword(runesWithSymbols)
}

// composePassword of passwordLength by selecting
// random elements from an ASCII subset (runePool).
func (p *Password) composePassword(runePool []rune) ([]rune, error) {
	var password []rune
	for i := int64(0); i < int64(p.passwordLength.value); i++ {
		runeAtIndex, err := p.randomIndexFromRunePool(runePool)
		if err != nil {
			return nil, ErrPassword
		}
		password = append(password, runePool[runeAtIndex])
	}
	return password, nil
}

// randomIndexFromRunePool returns index from Rune Pool using the
// Int() method from the Entropy implementation defined in the environment
func (p *Password) randomIndexFromRunePool(runePool []rune) (int64, error) {
	return p.Int(len(runePool) - 1)
}
