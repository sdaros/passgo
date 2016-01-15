package cli

import (
	"flag"
	"fmt"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
	"os"
	"reflect"
)

func ParseArgs(env *environment.Env) (cmd.CommandResult, error) {
	if err := parseOptions(env); err != nil {
		return nil, err
	}
	result, err := parseCommands(env)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseOptions(env *environment.Env) (err error) {
	env.Register("commandOptions", cmd.PassgoCommandOptions)
	optionsToParse := env.Lookup("commandOptions")
	for _, option := range optionsToParse.([]cmd.CommandOption) {
		name := reflect.ValueOf(option).Elem().FieldByName("name").String()
		description := reflect.ValueOf(option).Elem().FieldByName("description").String()
		flag.Var(option, name, description)
		env.Register(name, option)
	}
	flag.Parse()
	return nil
}

func parseCommands(env *environment.Env) (cmd.CommandResult, error) {
	if len(flag.Args()) == 0 {
		Usage()
		return nil, nil
	}
	// - flag.Args()[0] in env.Lookup("commands")
	if command, ok := cmd.PassgoCommands[flag.Args()[0]]; ok {
		result, err := command.Execute(env)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	Usage()
	return nil, nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
