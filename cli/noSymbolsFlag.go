package cli

import (
	"github.com/sdaros/passgo/cmd/password"
)

func NewNoSymbolsFlag() *password.NoSymbols {
	return password.NewNoSymbols()
}
