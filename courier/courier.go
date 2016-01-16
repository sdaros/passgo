// Package courier Proxies all requests from the client
// for example, seal() and stamp() ...
package courier

import (
	"github.com/sdaros/passgo/cmd"
	"github.com/sdaros/passgo/environment"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

// Courier acts as a CommandHandler and oversees program flow
// throughout the application.
type Courier struct {
	execute ExecuteFn
}

// ExecuteFn holds the Execute() method from an Executable Command.
type ExecuteFn func(*environment.Env) (cmd.CommandResult, error)

func (c *Courier) ProcessUserInput() {

}
func readFromStdIn() ([]byte, error) {
	fd := int(os.Stdin.Fd())
	pass, err := terminal.ReadPassword(fd)
	if err != nil {
		return nil, err
	}
	return pass, nil
}
