package cmd

// options defines all command line flags and validates input
// by implementing the flag.Value interface.

import (
	"flag"
)

// Option associated with a Command.
type CommandOption interface {
	flag.Value
}

var PassgoCommandOptions = []CommandOption{
	NewPasswordLength(),
	NewNoSymbols(),
}
