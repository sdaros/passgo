package cmd

import (
	"errors"

	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cmd/generate"
	"github.com/sdaros/passgo/cmd/password"
)

var (
	ErrGenerate = errors.New("cmd: Error while trying" +
		"to generate a new secret")
)

// Generate creates a new secret by taking the provided Url and Username and
// appending a randomly generated password.
type Generate struct {
	*app.App
	execute        func() (CmdResult, error)
	name           string
	noSymbols      *password.NoSymbols
	passwordLength *password.Length
	result         CmdResult
	url            *generate.Url
	userName       *generate.UserName
}

// NewGenerate returns a secret with default values
func NewGenerate() *Generate {
	generate := &Generate{
		App:            app.Null(),
		name:           "generate",
		noSymbols:      password.NewNoSymbols(),
		passwordLength: password.NewLength(),
		url:            generate.NewUrl(),
		userName:       generate.NewUserName(),
	}
	generate.execute = generateExecuteFn(generate)
	return generate
}

// generateExecuteFn validates command options then returns a
// an Envelope with the sealed Secret.
func generateExecuteFn(g *Generate) func() (CmdResult, error) {
	generateFn := func() (CmdResult, error) {
		g.ApplyCommandFlagsFrom(g.App)
		if err := g.validate(); err != nil {
			return nil, err
		}
		passwordSubCommand := g.executeSubCommands()[0]
		passwordCmdResult, err := passwordSubCommand()
		if err != nil {
			return nil, err
		}
		g.result = passwordCmdResult
		return g.result, nil
	}
	return generateFn
}

// ExecuteFn return an Envelope with the sealed Secret.
func (g *Generate) ExecuteFn() func() (CmdResult, error) { return g.execute }

// executeSubCommands executes dependencies (subcommands) required by Generate.
// Generate is dependent on only the Password subcommand.
func (g *Generate) executeSubCommands() [1]func() (CmdResult, error) {
	var executeSubCommandFuncs [1]func() (CmdResult, error)
	passwordSubCommand := NewPassword()
	passwordSubCommand.App = g.App
	executeSubCommandFuncs[0] = passwordSubCommand.ExecuteFn()
	return executeSubCommandFuncs
}

func (g *Generate) ApplyCommandFlagsFrom(passgo *app.App) error {
	if passgo == nil {
		return errors.New("We need a valid Passgo object to retrieve flags")
	}
	g.App = passgo
	if g.App.Lookup("user-name") != nil {
		unFromApp := g.App.Lookup("user-name").(*generate.UserName)
		g.userName = unFromApp
	} // else, user-name param was not provided; so the default will be used.
	if g.App.Lookup("url") != nil {
		urlFlag := g.App.Lookup("url").(*generate.Url)
		g.url = urlFlag
	} // else, url param was not provided; so the default will be used.
	if g.App.Lookup("password-length") != nil {
		passwordLengthFlag := g.App.Lookup("password-length").(*password.Length)
		g.passwordLength = passwordLengthFlag
	} // else, password-length param was not provided; so the default will be used.
	if g.App.Lookup("no-symbols") != nil {
		noSymbolsFlag := g.App.Lookup("no-symbols").(*password.NoSymbols)
		g.noSymbols = noSymbolsFlag
	} // else, no-symbols param was not provided; so the default will be used.
	return nil
}

func (g *Generate) validate() (err error) {
	if err := g.userName.Validate(nil); err != nil {
		return err
	}
	if err := g.url.Validate(nil); err != nil {
		return err
	}
	return nil
}

func (g *Generate) Name() string { return g.name }
