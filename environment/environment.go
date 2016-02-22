package environment

import (
	ent "github.com/sdaros/passgo/entropy"
	"github.com/sdaros/passgo/sealer"
	"github.com/sdaros/passgo/stamper"
)

// Env is provided as an environment by that is accessible
// to all clients that require its functionality.
type Env struct {
	Logger
	ent.Entropy
	sealer.
}

// Initialise the environment.
func Environment(logger Logger, entropy ent.Entropy) *Env {
	// nil logger does nothing
	if logger == nil {
		logger = new(NullLogger)
	}
	// nil entropy defaults to CryptoRand implementation
	if entropy == nil {
		entropy = ent.CryptoRand
	}

	return &Env{Logger: logger, Entropy: entropy}
}

func Null() *Env {
	return Environment(nil, nil)
}
