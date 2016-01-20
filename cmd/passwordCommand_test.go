package cmd

import (
	"fmt"
	"testing"
	"unicode"
)

type testVector string

// testVector must implement fmt.Stringer
func (tv testVector) String() string {
	return fmt.Sprintf("%v", string(tv))
}

func Test_password_against_invalid_password_length(t *testing.T) {
	command := NewPassword()

	// passwordLengthFlag too short
	command.passwordLength.value = 0
	_, err := command.Execute()
	if err == nil {
		t.Error("Should have received an error")
	}
	// passwordLengthFlag too long
	command.passwordLength.value = 257
	_, err = command.Execute()
	if err == nil {
		t.Error("Should have received an error")
	}
}

func Test_password_matches_a_provided_length(t *testing.T) {
	command := NewPassword()

	command.passwordLength.value = 256
	command.noSymbols.value = false
	cmdResult, err := command.Execute()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if cmdResult == nil || len(cmdResult.Value.(string)) != 256 {
		t.Errorf("Expected positive non nil length of password")
	}
}

func Test_password_with_and_without_symbols(t *testing.T) {
	command := NewPassword()

	// password should be generated with symbols
	command.passwordLength.value = 50
	command.noSymbols.value = false
	cmdResult, err := command.Execute()
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
	command.passwordLength.value = 50
	command.noSymbols.value = true
	cmdResult, err = command.Execute()
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

func Test_password_with_default_flags(t *testing.T) {
	command := NewPassword()
	_, err := command.Execute()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
