// Package courier Proxies all requests from the client
// for example, seal() and stamp() ...
package courier

import (
	"github.com/sdaros/passgo/cli"
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

// Courier acts as a CommandHandler and oversees program flow
// throughout the application.
type Courier struct {
	execute cmd.ExecuteFn
}

func (c *Courier) ProcessUserInput(env *environment.Env) (cmd.CommandResult, error) {
	commandToExecute, err := cli.Parse(env)
	if err != nil {
		return err
	}

}
func readFromStdIn() ([]byte, error) {
	fd := int(os.Stdin.Fd())
	pass, err := terminal.ReadPassword(fd)
	if err != nil {
		return nil, err
	}
	return pass, nil
}
