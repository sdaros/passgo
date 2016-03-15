package generate

import (
	"fmt"

	"github.com/sdaros/passgo/courier"
)

type UserName struct {
	name  string
	usage string
	value string
}

func NewUserName() *UserName {
	return &UserName{
		name:      "user-name",
		usage:     "Username associated with a secret.",
		isCommand: false,
	}
}

func (un *UserName) Name() string {
	return un.name
}

func (un *UserName) Usage() string {
	return un.usage
}

func (un *UserName) Value() string {
	return un.value
}

func (un *UserName) String() string {
	return fmt.Sprint(un.value)
}

func (un *UserName) Set(value string) (err error) {
	if err := un.Validate(value); err != nil {
		return err
	}
	un.value = value
	return nil
}

func (un *UserName) Validate(value interface{}) (err error) {
	if value != nil {
		return courier.IsValidUtf8String(value.(string))
	}
	// Explicitly validate UserName.value when called with nil.
	return courier.IsValidUtf8String(un.value)
}
