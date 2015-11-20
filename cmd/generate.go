package cmd

import (
  "github.com/sdaros/passgo/entropy"
  "fmt"
)

const (
  passwordLengthMin = 1
  passwordLengthMax = 256
)
// Generate returns a password to the caller
// based on the options provided to it.
type Generate struct {
  NoSymbols bool
  PasswordLength int
  EntropyImplementation entropy.Entropy
}

func NewGenerate() (*Generate) {
  // set defaults
  noSymbols := false
  passwordLength := 15
  entropyImplementation := entropy.CryptoRand
  args := &Generate{noSymbols, passwordLength, entropyImplementation}
  return args
}

func (args *Generate) Execute() (password []rune, err error) {
  if err := args.validateParameters(); err != nil {
    return nil, err
  }
  password, err = args.EntropyImplementation.Password(args)
  if err != nil {
    return nil, err
  }
  return password, nil
}

func (args *Generate) validateParameters() (error) {
  if args.PasswordLength < passwordLengthMin || args.PasswordLength > passwordLengthMax {
    err := fmt.Errorf( "Password length must be between %v and %v characters",
      passwordLengthMin, passwordLengthMax)
    return err
  }
  return nil
}
