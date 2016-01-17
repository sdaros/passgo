package cli

// params defines all command line flags and validates input
// by implementing the flag.Value interface.

import (
	"flag"
	"github.com/sdaros/passgo/cmd"
)

// Param that is given on the command line.
type PassgoFlag interface {
	flag.Value
	Name() string
	Usage() string
	IsCommand() bool
}

var passgoFlags = []PassgoFlag{
	cmd.NewPasswordCommandFlag(),
	cmd.NewNoSymbolsFlag(),
	cmd.NewPasswordLengthFlag(),
	cmd.NewGenerateCommandFlag(),
}
