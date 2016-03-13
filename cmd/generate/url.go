package generate

import (
	"fmt"

	"github.com/sdaros/passgo/courier"
)

type Url struct {
	name      string
	usage     string
	value     string
	isCommand bool
}

// NewUrl returns a new empty UserNameFlag.
func NewUrl() *Url {
	return &Url{
		name:      "Url",
		usage:     "Url associated with a secret.",
		isCommand: false,
	}
}

func (ur *Url) Name() string {
	return ur.name
}

func (ur *Url) Usage() string {
	return ur.usage
}

func (ur *Url) IsCommand() bool {
	return ur.isCommand
}

// String is provided to satisfy flag.Value interface.
func (ur *Url) String() string {
	return fmt.Sprint(ur.value)
}

// Set sets the value for the Url and validates it.
func (ur *Url) Set(fromCli string) (err error) {
	if err := ur.Validate(fromCli); err != nil {
		return err
	}
	ur.value = fromCli
	return nil
}

func (ur *Url) Validate(value interface{}) (err error) {
	if value != nil {
		return courier.IsValidUtf8String(value.(string))
	}
	// Explicitly validate Url.value when called with nil.
	return courier.IsValidUtf8String(ur.value)
}
