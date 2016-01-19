package cmd

import (
	"fmt"
	_ "github.com/sdaros/passgo/entropy"
	"github.com/sdaros/passgo/environment"
	"testing"
	"unicode"
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
	passwordLengthFlag := NewPasswordLengthFlag()
	noSymbolsFlag := NewNoSymbolsFlag()
	env.Register("password-length", passwordLengthFlag)
	env.Register("no-symbols", noSymbolsFlag)

	// passwordLengthFlag too short
	passwordLengthFlag.value = 0
	_, err := command.Execute(env)
	if err == nil {
		t.Error("Should have received an error")
	}
	// passwordLengthFlag too long
	passwordLengthFlag.value = 257
	env.Register("password-length", passwordLengthFlag)
	_, err = command.Execute(env)
	if err == nil {
		t.Error("Should have received an error")
	}
}

func Test_password_with_and_without_symbols(t *testing.T) {
	env := environment.Null()
	command := NewPassword()
	passwordLengthFlag := NewPasswordLengthFlag()
	noSymbolsFlag := NewNoSymbolsFlag()

	// password should be generated with symbols
	passwordLengthFlag.value = 50
	noSymbolsFlag.value = false
	env.Register("password-length", passwordLengthFlag)
	env.Register("no-symbols", noSymbolsFlag)
	cmdResult, err := command.Execute(env)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	isSymbolFound := false
	result := cmdResult.Value.(string)
	for _, r := range result {
		if unicode.IsSymbol(r) {
			isSymbolFound = true
			break
		}
	}
	if !isSymbolFound {
		t.Error("Expected to find a symbol in password")
	}
	// password should be generated without symbols
	passwordLengthFlag.value = 50
	noSymbolsFlag.value = true
	env.Register("password-length", passwordLengthFlag)
	env.Register("no-symbols", noSymbolsFlag)
	cmdResult, err = command.Execute(env)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	isSymbolFound = false
	result = cmdResult.Value.(string)
	for _, r := range result {
		if unicode.IsSymbol(r) {
			isSymbolFound = true
			break
		}
	}
	if isSymbolFound {
		t.Error("Expected not to find a symbol in password")
	}

}

func Test_password_matches_a_provided_length(t *testing.T) {
	env := environment.Null()
	command := NewPassword()
	passwordLengthFlag := NewPasswordLengthFlag()
	noSymbolsFlag := NewNoSymbolsFlag()

	passwordLengthFlag.value = 256
	noSymbolsFlag.value = false
	env.Register("password-length", passwordLengthFlag)
	env.Register("no-symbols", noSymbolsFlag)
	cmdResult, err := command.Execute(env)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if cmdResult == nil || len(cmdResult.Value.(string)) != 256 {
		t.Errorf("Expected positive non nil length of password")
	}
}
