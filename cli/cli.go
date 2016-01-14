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
		fmt.Println(reflect.TypeOf(command))
	}
	return nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
