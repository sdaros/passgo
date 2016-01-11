package courier

import (
	"flag"
	"fmt"
	"github.com/sdaros/passgo/cmd"
	"os"
)

type options struct {
	passwordLength        *int
	noSymbols             *bool
	entropyImplementation *string
}

func ParseOptions() error {
	opt := new(options)
	opt.passwordLength = flag.Int("password-length", 15,
		"Length of password to be generated")
	opt.noSymbols = flag.Bool("no-symbols", false,
		"Use only alphanumeric characters")
	opt.entropyImplementation = flag.String("entropy", "cryptoRand",
		"Specifies the entropy implementation to use")
	flag.Parse()
	if err := opt.executeCommands(); err != nil {
		return err
	}
	return nil
}

func (opt *options) executeCommands() error {
	if len(flag.Args()) == 0 {
		Usage()
		return nil
	}
	switch flag.Args()[0] {
	case "password":
		passwordCommand := new(cmd.Password)
		password, err := passwordCommand.Execute(opt.noSymbols, opt.passwordLength,
			opt.entropyImplementation)
		if err != nil {
			return err
		}
		fmt.Printf("Generated password: %v\n", string(password))
	default:
		Usage()
	}
	return nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
