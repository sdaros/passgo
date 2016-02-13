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
		// get sub
		g.result = &CmdResult{Value: "heelo"}
		return g.result, nil
	}
	return generateFn
}

// ExecuteFn return an Envelope with the sealed Secret.
func (g *Generate) ExecuteFn() func() (*CmdResult, error) { return g.execute }

func (g *Generate) test() []func() string {
	var fns []func() string
	fn := func() string { return "bar" }
	fns = append(fns, fn)
	return fns

}

// executeSubCommands executes dependencies (subcommands) required by Generate
func (g *Generate) executeSubCommands() []func() (*CmdResult, error) {
	var executeSubCommandsFuncs []func() (*CmdResult, error)
	passwordSubCommandFn := func() (*CmdResult, error) {
		p := NewPassword()
		// Apply command flags given to generate through app
		p.ApplyCommandFlags(g.App)
		// NEXT: Why does it think there aren't enough commands to return?
		return p.ExecuteFn()
	}
	executeSubCommandsFuncs = append(executeSubCommandsFuncs, passwordSubCommandFn)
	return executeSubCommandsFuncs
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
