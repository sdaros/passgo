package cli

// params defines all command line flags and validates input
// by implementing the flag.Value interface.

import (
	"flag"
	"github.com/sdaros/passgo/cmd"
)

// Param that is given on the command line.
type Param interface {
	flag.Value
	Name() string
	Description() string
	IsCommand() bool
}

var PassgoFlags = flag.NewFlagSet("passgoFlags", flag.ExitOnError)

var PassgoParams = []Param{
	cmd.NewPasswordLengthParam(),
	cmd.NewNoSymbolsParam(),
	cmd.NewPasswordParam(),
	cmd.NewGenerateParam(),
}
