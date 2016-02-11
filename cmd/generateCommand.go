package cmd

import (
	"errors"
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/mailbag"
)

var (
	ErrGenerate = errors.New("cmd: Error while trying" +
		"to generate a new secret")
)

// Generate creates a new secret by taking the provided Url and Username and
// appending a randomly generated password.
type Generate struct {
	name           string `schema.org: "/name"`
	userName       *userNameFlag
	url            *urlFlag
	execute        func() (*CmdResult, error)
	passwordLength *passwordLengthFlag
	noSymbols      *noSymbolsFlag
	result         *CmdResult
	*app.App
}

// NewGenerate returns a secret with default values
func NewGenerate() *Generate {
	generate := &Generate{
		name:           "generate",
		userName:       NewUserNameFlag(),
		url:            NewUrlFlag(),
		noSymbols:      NewNoSymbolsFlag(),
		passwordLength: NewPasswordLengthFlag(),
		App:            app.Null(),
	}
	generate.execute = generateExecuteFn(generate)
	return generate
}

// generateExecuteFn validates command options then returns a
// an Envelope with the sealed Secret.
func generateExecuteFn(g *Generate) func() (*CmdResult, error) {
	generateFn := func() (*CmdResult, error) {
		secret := new(mailbag.Secret)
		secret.SetUrl("foobar")
		g.userName.value = secret.UserName()
		g.url.value = secret.Url()
		g.result = &CmdResult{Value: "heelo"}
		return g.result, nil
	}
	return generateFn
}

func (g *Generate) ExecuteFn() func() (*CmdResult, error) { return g.execute }

// NEXT: responsible for executing Password() and doing
// something with the value.
// ExecuteSubCommands executes dependencies required by Generate
func (g *Generate) executeSubCommands() error {
	return nil
}

func (g *Generate) ApplyCommandFlags() {
	unFromFlag := g.App.Lookup("user-name").(*userNameFlag)
	urlFromFlag := g.App.Lookup("url").(*urlFlag)
	if unFromFlag != nil {
		g.userName = unFromFlag
	} // User name flag not provided; use default.
	if urlFromFlag != nil {
		g.url = urlFromFlag
	} // Url flag not provided; use default.
}
