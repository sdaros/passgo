package cmd

import (
  _ "github.com/sdaros/passgo/entropy"
)

// Generate returns a password to the user
// based on the options provided to it.
type Generate struct {
  noSymbols bool
  passwordLength int
  password []rune
}

func NewGenerate() (pass []rune) {
  // set defaults
  noSymbols := false
  passwordLength := 15
  var password []rune
  generated := &Generate{noSymbols, passwordLength, password}
  if noSymbols {
    generated.getPasswordNoSymbols()
    return generated.password
  }
  generated.getPasswordWithSymbols()
  return generated.password
}

func (gen *Generate) getPasswordWithSymbols() {
}
func (gen *Generate) getPasswordNoSymbols() {
}
