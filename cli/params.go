package cli

// params defines all command line flags and validates input
// by implementing the flag.Value interface.

import (
	"flag"
)

// Param that is given on the command line.
type Param interface {
	flag.Value
}

var PassgoParams = []Param{
	NewPasswordLength(),
	NewNoSymbols(),
	NewPasswordParam(),
	NewGenerateParam(),
}
