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
		secret := populateSecret(g)
		return &CmdResult{Value: secret.Password()}, nil
	}
	return generateFn
}

func populateSecret(g *Generate) *mailbag.Secret {
	secret := new(mailbag.Secret)
	secret.SetPassword(g.result.String())
	secret.SetUserName(g.userName.value)
	secret.SetUrl(g.url.value)
	secret.SetNote("")
	return secret
}

// ExecuteFn return an Envelope with the sealed Secret.
func (g *Generate) ExecuteFn() func() (*CmdResult, error) { return g.execute }

// executeSubCommands executes dependencies (subcommands) required by Generate.
// Generate is dependent on only the Password subcommand.
func (g *Generate) executeSubCommands() [1]func() (*CmdResult, error) {
	var executeSubCommandFuncs [1]func() (*CmdResult, error)
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
		unFromFlag := g.App.Lookup("user-name").(*userNameFlag)
		g.userName = unFromFlag
	} // else, user-name flag was not provided; so the default will be used.
	if g.App.Lookup("url") != nil {
		urlFlag := g.App.Lookup("url").(*urlFlag)
		g.url = urlFlag
	} // else, url flag was not provided; so the default will be used.
	if g.App.Lookup("password-length") != nil {
		passwordLengthFlag := g.App.Lookup("password-length").(*passwordLengthFlag)
		g.passwordLength = passwordLengthFlag
	} // else, password-length flag was not provided; so the default will be used.
	if g.App.Lookup("no-symbols") != nil {
		noSymbolsFlag := g.App.Lookup("no-symbols").(*noSymbolsFlag)
		g.noSymbols = noSymbolsFlag
	} // else, no-symbols flag was not provided; so the default will be used.
	return nil
}

func (g *Generate) validate() (err error) {
	userName := g.userName.value
	if err := g.userName.Validate(userName); err != nil {
		return err
	}
	url := g.url.value
	if err := g.url.Validate(url); err != nil {
		return err
	}
	return nil
}

func (g *Generate) Name() string { return g.name }
