package main

import (
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cli"
	"github.com/sdaros/passgo/environment"
)

func main() {
	env := environment.Environment(new(environment.StandardLogger), nil)
	passgo := app.Passgo(env, nil)
	processInput(passgo)
	passgo.Info("app output:\n", passgo)
}

func processInput(passgo *app.App) {
	cli.Parse(passgo)
}

func displayOutput(pa *app.App) {}
