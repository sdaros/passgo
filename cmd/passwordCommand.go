package cmd

import (
	"encoding/json"
	"errors"
	"github.com/sdaros/passgo/app"
	"strconv"
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
	// TODO: JSON stringify uses HTML_Escape for & < > etc.
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

// Password returns a password to the caller based on the parameters provided to it.
type Password struct {
	name           string
	execute        func() (string, error)
	noSymbols      *noSymbolsFlag
	passwordLength *passwordLengthFlag
	result         string
	*app.App
}

// NewPassword returns a password command with default values
func NewPassword() *Password {
	password := &Password{
		name:           "password",
		noSymbols:      NewNoSymbolsFlag(),
		passwordLength: NewPasswordLengthFlag(),
		App:            app.Null(),
	}
	password.execute = passwordExecuteFn(password)
	return password
}

// passwordExecuteFn validates command options then returns a password
// composed of random elements chosen from a rune pool
func passwordExecuteFn(p *Password) func() (string, error) {
	passwordExecuteFn := func() (string, error) {
		if err := p.validate(); err != nil {
			return "", err
		}
		if p.noSymbols.value {
			result, err := p.composePassword(runesNoSymbols)
			if err != nil {
				return "", err
			}
			p.result = result
			return p.String()
		}
		result, err := p.composePassword(runesWithSymbols)
		if err != nil {
			return "", err
		}
		p.result = result
		return p.String()

	}
	return passwordExecuteFn
}

func (p *Password) ExecuteFn() func() (string, error) { return p.execute }

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

func (p *Password) ApplyCommandFlags(passgo *app.App) {
	plFromFlag := passgo.Lookup("password-length").(*passwordLengthFlag)
	nsFromFlag := passgo.Lookup("no-symbols").(*noSymbolsFlag)
	if plFromFlag != nil {
		p.passwordLength = plFromFlag
	} // else, Password length flag not provided; use default.
	if nsFromFlag != nil {
		p.noSymbols = nsFromFlag
	} // else, No symbols flag not provided; use default.
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

func (p *Password) String() (string, error) {
	type PasswordResult struct {
		Name           string
		NoSymbols      string
		PasswordLength int
		Result         string
	}
	passwordLength, err := strconv.Atoi(p.passwordLength.String())
	if err != nil {
		return "", nil
	}
	pResult := &PasswordResult{
		p.name,
		p.noSymbols.String(),
		passwordLength,
		p.result,
	}
	jsonResult, err := json.MarshalIndent(pResult, " ", "\t")
	if err != nil {
		return "", err
	}
	return string(jsonResult), nil

}
