package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cli"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
)

const (
	// buildMetadata is replaced when package is built using -ldflags -X
	// ex: go build -ldflags "-X main.buildMetadata=`git rev-parse HEAD`"
	buildMetadata = "<placeholder>"
	version       = "0.1.0"
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
	if err := command.ApplyCommandParamsFrom(passgo); err != nil {
		return err
	}
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

func Version() string { return version + "+" + buildMetadata }
