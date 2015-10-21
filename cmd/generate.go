package cmd

import (
  "github.com/sdaros/passgo/entropy"
)

// Generate returns a password to the user
// based on the options provided to it.
type Generate struct {
  noSymbols bool
  passwordLength int
  entropyImplementation entropy.Entropy
}

func NewGenerate() (*Generate) {
  // set defaults
  noSymbols := false
  passwordLength := 15
  entropyImplementation := entropy.CryptoRand
  args := &Generate{noSymbols, passwordLength, entropyImplementation}
  //password = entropy.Password(args)
  //TODO: that should be done in courier
  return args
}
