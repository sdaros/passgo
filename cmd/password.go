package cmd

import (
  "github.com/sdaros/passgo/entropy"
_  "reflect"
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
  PasswordLength int64
  EntropyImplementation entropy.Entropy
}

// Execute validates parameters of Password and runs the Password Command.
func (p *Password) Execute(options ...interface{}) (password []rune, err error) {
  if err = p.validate(options...); err != nil {
    return nil, err
  }
  password, err = p.Password()
  if err != nil {
    return nil, err
  }
  return password, nil
}

// Password returns a password by selecting random
// elements from an ASCII subset (runePool).
func (p *Password) Password() (password []rune, err error) {
  if p.NoSymbols {
    return p.composePassword(runesNoSymbols)
  }
  return p.composePassword(runesWithSymbols)
}

// validate validates the options for the Password Command against predefined rules.
func (p *Password) validate(options ...interface{}) error {
  noSymbols := *options[0].(*bool)
  passwordLength := *options[1].(*int)
  entropyImplementation := *options[2].(*string)
  if passwordLength < passwordLengthMin || passwordLength > passwordLengthMax {
    err := fmt.Errorf( "Password length must be between %v and %v characters",
    passwordLengthMin, passwordLengthMax)
    return err
  }
  switch entropyImplementation {
  case "cryptoRand":
    p.EntropyImplementation = entropy.CryptoRand
  default:
    err := fmt.Errorf("the specified entropy implementation, %v does not exist",
    entropyImplementation)
    return err
  }
  p.PasswordLength = int64(passwordLength)
  p.NoSymbols = noSymbols
  return nil
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
