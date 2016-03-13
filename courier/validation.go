package courier

import (
	"fmt"
	"unicode/utf8"
)

func IsValidUtf8String(value string) error {
	if !utf8.ValidString(value) {
		return fmt.Errorf("Expected username to be a valid utf8 string,"+
			" got %q instead", value)
	}
	return nil
}

func IsBetween(min, max, value int) error {
	if value < min || value > max {
		err := fmt.Errorf("Password length must be between %v and %v characters",
			min, max)
		return err
	}
	return nil
}
