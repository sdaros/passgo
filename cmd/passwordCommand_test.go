package cmd

import (
	"fmt"
	"testing"
	"unicode"
)

// NEXT: refactor test to use table
// see github.com/golang/go/wiki/TableDrivenTests
type testVector string

// testVector must implement fmt.Stringer
func (tv testVector) String() string {
	return fmt.Sprintf("%v", string(tv))
}

func Test_password_against_invalid_password_length(t *testing.T) {
	command := NewPassword()
	exec := command.execute

	// passwordLengthFlag too short
	command.passwordLength.value = 0
	_, err := exec()
	if err == nil {
		t.Error("Should have received an error")
	}
	// passwordLengthFlag too long
	command.passwordLength.value = 257
	_, err = exec()
	if err == nil {
		t.Error("Should have received an error")
	}
}

func Test_password_matches_a_provided_length(t *testing.T) {
	command := NewPassword()
	exec := command.execute

	command.passwordLength.value = 256
	command.noSymbols.value = false
	result, err := exec()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.String() == "" || len(result.String()) != 256 {
		t.Errorf("Expected positive non nil length of password")
	}
}

func Test_password_with_and_without_symbols(t *testing.T) {
	command := NewPassword()
	exec := command.execute

	// password should be generated with symbols
	command.passwordLength.value = 256
	command.noSymbols.value = false
	result, err := exec()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	isSymbolFound := false
	for _, r := range result.String() {
		if unicode.IsSymbol(r) {
			isSymbolFound = true
			break
		}
	}
	if !isSymbolFound {
		t.Error("Expected to find a symbol in password")
	}
	// password should be generated without symbols
	command.passwordLength.value = 256
	command.noSymbols.value = true
	result, err = exec()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	isSymbolFound = false
	for _, r := range result.String() {
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
	exec := command.execute
	_, err := exec()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
