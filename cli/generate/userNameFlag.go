package cli

import (
	"github.com/sdaros/passgo/cmd/generate"
)

func NewUserNameFlag() *generate.UserName {
	return generate.NewUserName()
}
