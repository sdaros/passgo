package generate

import (
	"fmt"

	"github.com/sdaros/passgo/courier"
)

type Url struct {
	name  string
	usage string
	value string
}

// NewUrl returns a new empty UserNameFlag.
func NewUrl() *Url {
	return &Url{
		name:  "url",
		usage: "Url associated with a secret.",
	}
}

func (ur *Url) Name() string {
	return ur.name
}

func (ur *Url) Usage() string {
	return ur.usage
}

func (ur *Url) Value() string {
	return ur.value
}

// String is provided to satisfy flag.Value interface.
func (ur *Url) String() string {
	return fmt.Sprint(ur.value)
}

// Set sets the value for the Url and validates it.
func (ur *Url) Set(value string) (err error) {
	if err := ur.Validate(value); err != nil {
		return err
	}
	ur.value = value
	return nil
}

func (ur *Url) Validate(value interface{}) (err error) {
	if value != nil {
		return courier.IsValidUtf8String(value.(string))
	}
	// Explicitly validate Url.value when called with nil.
	return courier.IsValidUtf8String(ur.value)
}
