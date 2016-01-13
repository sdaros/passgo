package cmd

import (
	"flag"
	"fmt"
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
	// register options from `options.go` in the environment
	registerOptions(env)
	registeredOptions := env.Lookup("options")
	optionsToParse := registeredOptions.(options)
	for _, option := range optionsToParse {
		name := reflect.ValueOf(option).Elem().FieldByName("name").String()
		description := reflect.ValueOf(option).Elem().FieldByName("description").String()
		flag.Var(option.(flag.Value), name, description)
	}
	flag.Parse()
	return nil
}

func parseCommands(env *environment.Env) (err error) {
	//registerCommands(env)
	return nil
}

func executeCommands(env *environment.Env) (err error) {
	/*	if len(flag.Args()) == 0 {
			Usage()
			return nil
		}
		// inject the environment object
		switch flag.Args()[0] {
		case "password":
			passwordCommand := &cmd.Password{opt.noSymbolsFlag, opt.passwordLengthFlag, opt.env}
			password, err := passwordCommand.Execute(opt.noSymbolsFlag, opt.passwordLengthFlag)
			if err != nil {
				return err
			}
			fmt.Printf("Generated password: %v\n", string(password))
		default:
			Usage()
		}*/
	return nil
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}
