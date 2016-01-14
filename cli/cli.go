package cli

import (
	"flag"
	"fmt"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
	"os"
	"reflect"
)

func ParseArgs(env *environment.Env) (err error) {
	if err := parseOptions(env); err != nil {
		return err
	}
	if err := parseCommands(env); err != nil {
		return err
	}
	return nil
}

func parseOptions(env *environment.Env) (err error) {
	cmd.RegisterOptions(env)
	registeredOptions := env.Lookup("options")
	optionsToParse := registeredOptions.(cmd.Options)
	for _, option := range optionsToParse {
		name := reflect.ValueOf(option).Elem().FieldByName("name").String()
		description := reflect.ValueOf(option).Elem().FieldByName("description").String()
		flag.Var(option.(flag.Value), name, description)
	}
	flag.Parse()
	return nil
}

func parseCommands(env *environment.Env) (err error) {
	cmd.RegisterCommands(env)
	if err := executeCommands(env); err != nil {
		return err
	}
	return nil
}

func executeCommands(env *environment.Env) (err error) {
	if len(flag.Args()) == 0 {
		Usage()
		return nil
	}
	// - flag.Args()[0] in env.Lookup("commands")
	registeredCommands := env.Lookup("commands")
	commandsToParse := registeredCommands.(cmd.Commands)
	for _, command := range commandsToParse {
		commandName := reflect.ValueOf(command).Elem().FieldByName("name").String()
		if flag.Args()[0] == commandName {
			// - use docker cli/cli.go for inspiration
			// probably something like this:
			// - method := reflect.ValueOf(command).MethodByName(execute)
			// - return method.Interface().(func(...string) error), nil
			fmt.Println(commandName)
			return nil

		}
	}
	return nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
