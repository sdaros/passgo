package cmd

import (
	"fmt"
	"github.com/sdaros/passgo/entropy"
	"github.com/sdaros/passgo/environment"
	"testing"
)

type testVector string

// testVector must implement fmt.Stringer
func (tv testVector) String() string {
	return fmt.Sprintf("%v", string(tv))
}

// Password() should failed
func Test_password_against_invalid_password_length(t *testing.T) {
	env := environment.Null()
	cmdTooShort := &Password{false, 0, entropy.CryptoRand}
	_, err := cmdTooShort.Execute()
	if err == nil {
		t.Error("Should have received an error.")
	}
	cmdTooLong := &Password{false, 257, entropy.CryptoRand}
	_, err = cmdTooLong.Execute()
	if err == nil {
		t.Error("Should have received an error.")
	}
}
