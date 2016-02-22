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
	passgo.Register("passgoFlags", cmd.PassgoFlags)
	flagsToParse := passgo.Lookup("passgoFlags").([]cmd.PassgoFlag)
	for _, flag := range flagsToParse {
		flagSet.Var(flag, flag.Name(), flag.Usage())
		passgo.Register(flag.Name(), flag)
	}
	flagSet.Parse(os.Args[1:])
	flagSet.Visit(thenRegisterCommandToExecute(passgo))
	// No command was provided by the user on the command line; print Usage.
	if passgo.Lookup("commandToExecute") == nil {
		flagSet.Usage()
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
