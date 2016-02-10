package main

import (
	"fmt"
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cli"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
)

func main() {
	env := environment.Environment(new(environment.StandardLogger), nil)
	passgo := app.Passgo(env, nil)
	result, err := processInput(passgo)
	if err != nil {
		passgo.Info("error:\n", err)
	}
	fmt.Printf("app output:\n", result)
}

func processInput(passgo *app.App) (result string, err error) {
	cli.Parse(passgo)
	command := passgo.Lookup("commandToExecute").(cmd.Command)
	execute := command.ExecuteFn()
	result, err = execute()
	if err != nil {
		return "", err
	}
	return result, nil

}

func displayOutput(pa *app.App) {}
