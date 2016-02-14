package cmd

import (
	"flag"
)

// PassgoFlag is parsed when executing commands.
type PassgoFlag interface {
	flag.Value
	Name() string
	Usage() string
	IsCommand() bool
}

var PassgoFlags = []PassgoFlag{
	NewPasswordCommandFlag(),
	NewNoSymbolsFlag(),
	NewUrlFlag(),
	NewUserNameFlag(),
	NewPasswordLengthFlag(),
	NewGenerateCommandFlag(),
}
