package cli

import (
	"flag"
	"fmt"
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cmd"
	"os"
)

func Parse(passgo *app.App) {
	flagSet := flag.NewFlagSet("passgoFlags", flag.ExitOnError)
	setUsage(flagSet)
	registerCliFlagsWithPassgoRegistrar(passgo, flagSet)
	flagSet.Parse(os.Args[1:])
	flagSet.Visit(thenRegisterCommandToExecute(passgo))
	if passgo.Lookup("commandToExecute") == nil {
		// No executable command was provided by the user
		// on the command line; print Usage.
		flagSet.Usage()
	}
}

func registerCliFlagsWithPassgoRegistrar(passgo *app.App, flagSet *flag.FlagSet) {
	passgo.Register("passgoFlags", cmd.PassgoFlags)
	flagsToParse := passgo.Lookup("passgoFlags").([]cmd.PassgoFlag)
	for _, flag := range flagsToParse {
		flagSet.Var(flag, flag.Name(), flag.Usage())
		passgo.Register(flag.Name(), flag)
	}
}

func thenRegisterCommandToExecute(passgo *app.App) func(*flag.Flag) {
	fn := func(f *flag.Flag) {
		currentFlag := f.Value.(cmd.PassgoFlag)
		if currentFlag.IsCommand() {
			commandToExecute := cmd.PassgoCommands[currentFlag.Name()]
			passgo.Register("commandToExecute", commandToExecute)
		}
	}
	return fn
}

func setUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fs.PrintDefaults()
		os.Exit(2)
	}
}
