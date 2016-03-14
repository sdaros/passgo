package cli

import (
	"github.com/sdaros/passgo/cmd/password"
)

func NewPasswordLengthFlag() *password.Length {
	return password.NewLength()
}
