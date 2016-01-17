package cli

import (
	"flag"
	"fmt"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
	"os"
)

func Parse(env *environment.Env) (cmd.ExecuteFn, error) {
	var commandsToExecute []Param
	env.Register("params", PassgoParams)
	paramsToParse := env.Lookup("params")
	for _, param := range paramsToParse.([]Param) {
		if param.IsCommand() {
			commandsToExecute = append(commandsToExecute, param)
		}
		flag.Var(param, param.Name(), param.Description())
	}
	flag.Parse()

	command := cmd.PassgoCommands[commandsToExecute[0].Name()]
	return command.Execute, nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
