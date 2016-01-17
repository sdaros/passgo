package cmd

import (
	"fmt"
	_ "github.com/sdaros/passgo/entropy"
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
	command := NewPassword()
	passwordLength := NewPasswordLengthFlag()

	// passwordLength too short
	passwordLength.value = 0
	env.Register("password-length", passwordLength)
	_, err := command.Execute(env)
	if err == nil {
		t.Error("Should have received an error")
	}
	// passwordLength too long
	passwordLength.value = 257
	env.Register("password-length", passwordLength)
	_, err = command.Execute(env)
	if err == nil {
		t.Error("Should have received an error")
	}
}
