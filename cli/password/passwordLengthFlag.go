package cli

import (
	"github.com/sdaros/passgo/cmd"
)

func NewPasswordLengthFlag() *cmd.PasswordLength {
	return cmd.NewPasswordLength()
}
