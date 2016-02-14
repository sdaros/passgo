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
		if err := g.validate(); err != nil {
			return nil, err
		}
		g.ApplyCommandFlags(g.App)
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
	// Apply command flags given to generate through app
	passwordSubCommand.ApplyCommandFlags(g.App)
	executeSubCommandFuncs[0] = passwordSubCommand.ExecuteFn()
	return executeSubCommandFuncs
}

func (g *Generate) ApplyCommandFlags(passgo *app.App) {
	unFromFlag := passgo.Lookup("user-name").(*userNameFlag)
	urlFromFlag := passgo.Lookup("url").(*urlFlag)
	plFromFlag := passgo.Lookup("password-length").(*passwordLengthFlag)
	nsFromFlag := passgo.Lookup("no-symbols").(*noSymbolsFlag)
	if unFromFlag != nil {
		g.userName = unFromFlag
	}
	if urlFromFlag != nil {
		g.url = urlFromFlag
	}
	if plFromFlag != nil {
		g.passwordLength = plFromFlag
	}
	if nsFromFlag != nil {
		g.noSymbols = nsFromFlag
	}
	g.App = passgo
}

func (g *Generate) validate() (err error) {
	// TODO: implement validation
	return nil
}

func (g *Generate) Name() string { return g.name }
