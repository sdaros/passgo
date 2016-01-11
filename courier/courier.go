// Package Courier Proxies all requests from the client
// for example, seal() and stamp() ...
package courier

import (
	_ "github.com/sdaros/passgo/cmd"
	ent "github.com/sdaros/passgo/entropy"
	_ "github.com/sdaros/passgo/sealer"
	_ "github.com/sdaros/passgo/stamper"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func readFromStdIn() ([]byte, error) {
	fd := int(os.Stdin.Fd())
	pass, err := terminal.ReadPassword(fd)
	if err != nil {
		return nil, err
	}
	return pass, nil
}

// Env is provided as an environment by courier that is accessible
// to all clients that require its functionality
type Env struct {
	Logger
	ent.Entropy
}

// Initialise the environment
func Environment(logger Logger, entropy ent.Entropy) *Env {
	// TODO: Read from config file

	// nil logger does nothing
	if logger == nil {
		logger = new(NullLogger)
	}
	// nil entropy default to CryptoRand implementation
	if entropy == nil {
		entropy = ent.CryptoRand
	}
	return &Env{Logger: logger, Entropy: entropy}
}
