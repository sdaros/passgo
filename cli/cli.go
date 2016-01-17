package cli

import (
	"flag"
	"fmt"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
	"os"
)

func Parse(env *environment.Env) {
	flagSet := flag.NewFlagSet("passgoFlags", flag.ExitOnError)
	setUsage(flagSet)
	env.Register("passgoFlags", passgoFlags)
	flagsToParse := env.Lookup("passgoFlags").([]PassgoFlag)
	for _, flag := range flagsToParse {
		flagSet.Var(flag, flag.Name(), flag.Usage())
		env.Register(flag.Name(), flag)
	}
	flagSet.Parse(os.Args[1:])
	flagSet.Visit(toRegisterCommandInEnv(env))
	if env.Lookup("commandToExecute") == nil {
		flagSet.Usage()
	}
}

func toRegisterCommandInEnv(env *environment.Env) func(*flag.Flag) {
	fn := func(f *flag.Flag) {
		if f.Value.(PassgoFlag).IsCommand() {
			commandToExecute := cmd.PassgoCommands[f.Value.(PassgoFlag).Name()]
			env.Register("commandToExecute", commandToExecute)
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
