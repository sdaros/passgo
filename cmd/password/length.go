package password

import (
	"fmt"
	"strconv"

	"github.com/sdaros/passgo/courier"
)

type Length struct {
	name  string
	usage string
	value int
}

const (
	lengthMin = 1
	lengthMax = 256
)

// NewPasswordLength returns a Length parameter with default values.
func NewLength() *Length {
	return &Length{
		name:  "password-length",
		usage: "Length of password to be generated.",
		value: 15,
	}
}

func (l *Length) Name() string {
	return l.name
}

func (l *Length) Usage() string {
	return l.usage
}

// String is provided to satisfy flag.Value interface.
func (l *Length) String() string {
	return fmt.Sprint(l.value)
}

func (l *Length) Value() int { return l.value }

func (l *Length) Validate(value interface{}) error {
	if value != nil {
		return courier.IsBetween(lengthMin, lengthMax, value.(int))
	}
	// Explicitly validate Length.value when called with nil.
	return courier.IsBetween(lengthMin, lengthMax, l.value)
}

// Set sets the value for the length and validates the range.
func (l *Length) Set(value string) (err error) {
	length, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := l.Validate(length); err != nil {
		return err
	}
	l.value = length
	return nil
}
