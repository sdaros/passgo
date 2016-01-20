package main

import (
	"github.com/sdaros/passgo/courier"
	"github.com/sdaros/passgo/environment"
)

func main() {
	env := environment.Environment(new(environment.StandardLogger), nil, nil)
	cr := new(courier.Courier)
	if err := cr.ProcessUserInput(env); err != nil {
		env.Error(err)
	}
	result, err := cr.Execute()
	if err != nil {
		env.Error(err)
	}
	jsonResult, err := result.String()
	if err != nil {
		env.Error(err)
	}
	env.Info("courier output:\n", jsonResult)
}
