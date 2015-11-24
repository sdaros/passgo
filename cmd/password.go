package cmd

import (
  "github.com/sdaros/passgo/entropy"
  "errors"
  "fmt"
)

const (
  passwordLengthMin = 1
  passwordLengthMax = 256
)

var (
  // Define a range of characters that the Password() implementation
  // can choose from when generating random passwords.
  runesNoSymbols = []rune {
  '0','1','2','3','4','5','6','7','8','9','A','B',
  'C','D','E','F','G','H','I','J','K','L','M','N',
  'O','P','Q','R','S','T','U','V','W','X','Y','Z',
  'a','b','c','d','e','f','g','h','i','j','k','l',
  'm','n','o','p','q','r','s','t','u','v','w','x',
  'y','z'}
  runesWithSymbols = []rune {
  '!','"','#','$','%','&','\'','(',')','*','+',',',
  '-','.','/','0','1','2','3','4','5','6','7','8',
  '9',':',';','=','>','?','@','A','B','C','D','E',
  'F','G','H','I','J','K','L','M','N','O','P','Q',
  'R','S','T','U','V','W','X','Y','Z','[','\\',']',
  '^','_','`','a','b','c','d','e','f','g','h','i',
  'j','k','l','m','n','o','p','q','r','s','t','u',
  'v','w','x','y','z','{','|','}','~'}
  ErrPassword = errors.New("cryptoRand: Error while trying" +
    " to generate password")
)

// Password returns a password to the caller
// based on the parameters provided to it.
type Password struct {
  NoSymbols bool
  PasswordLength int
  EntropyImplementation entropy.Entropy
}

func NewPasswordWithDefaults() (*Password) {
  // set parameter defaults.
  NoSymbols := false
  PasswordLength := 15
  EntropyImplementation := entropy.CryptoRand
  passwordCommand := &Password{NoSymbols, PasswordLength, EntropyImplementation}
  return passwordCommand
}

// Execute valides parameters of Password
// and then runs the Password Command.
func (p *Password) Execute() (password []rune, err error) {
  if err := p.validateParameters(); err != nil {
    return nil, err
  }
  password, err = p.Password()
  if err != nil {
    return nil, err
  }
  return password, nil
}

// validateParameters validates the fields of the
// Password struct against predefined rules.
func (p *Password) validateParameters() (error) {
  if p.PasswordLength < passwordLengthMin || p.PasswordLength > passwordLengthMax {
    err := fmt.Errorf( "Password length must be between %v and %v characters",
      passwordLengthMin, passwordLengthMax)
    return err
  }
  return nil
}

// Password returns a password by selecting random
// elements from an ASCII subset (runePool).
func (p *Password) Password() (password []rune, err error) {
  if p.NoSymbols {
    return p.composePassword(runesNoSymbols)
  }
  return p.composePassword(runesWithSymbols)
}

// composePassword of passwordLength by selecting
// random elements from an ASCII subset (runePool).
func (p *Password) composePassword(runePool []rune) ([]rune, error) {
  var password []rune
  for i := int64(0); i < int64(p.PasswordLength); i++ {
    runeAtIndex, err := p.randomIndexFromRunePool(runePool)
    if err != nil {
      return nil, ErrPassword
    }
    password = append(password, runePool[runeAtIndex])
  }
  return password, nil
}

func (p *Password) randomIndexFromRunePool(runePool []rune) (int64, error) {
  return p.EntropyImplementation.Int(len(runePool) - 1)
}
