package main

import (
	"fmt"
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cli"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
	"io"
	"os"
)

func main() {
	env := environment.Environment(new(environment.StandardLogger), nil)
	passgo := app.Passgo(env, nil)
	if err := processInput(passgo); err != nil {
		passgo.Info("error:\n", err)
	}
}

func processInput(passgo *app.App) error {
	cli.Parse(passgo)
	command := passgo.Lookup("commandToExecute").(cmd.Command)
	execute := command.ExecuteFn()
	cmdResult, err := execute()
	if err != nil {
		return err
	}
	_, err = displayOutputTo(os.Stdout, cmdResult)
	if err != nil {
		return err
	}
	return nil
}

func displayOutputTo(quill io.Writer, cmdResult cmd.CmdResult) (n int, err error) {
	return fmt.Fprint(quill, cmdResult)
}
