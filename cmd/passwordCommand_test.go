package cmd

import (
	"fmt"
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
	command := NewPassword(env)
	passwordLength := NewPasswordLength()

	// passwordLength too short
	passwordLength.value = 0
	command.passwordLength = passwordLength
	pass, err := command.Execute()
	fmt.Printf("command: %#v", command.passwordLength)
	fmt.Printf("pass: %v", pass)
	if err == nil {
		t.Error("Should have received an error")
	}
	// passwordLength too long
	passwordLength.value = 257
	command.passwordLength = passwordLength
	_, err = command.Execute()
	if err == nil {
		t.Error("Should have received an error.")
	}
}
