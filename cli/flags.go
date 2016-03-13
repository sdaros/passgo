package cli

import (
	"flag"
)

// PassgoFlag is parsed when executing commands.
type PassgoFlag interface {
	IsCommand() bool
	Name() string
	Usage() string
	flag.Value
}

var PassgoFlags = []PassgoFlag{
	NewGenerateFlag(),
	NewNoSymbolsFlag(),
	NewPasswordFlag(),
	NewPasswordLengthFlag(),
	NewScryptFlag(),
	NewStampFlag(),
	NewUrlFlag(),
	NewUserNameFlag(),
}
