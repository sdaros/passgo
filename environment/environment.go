package environment

import (
	ent "github.com/sdaros/passgo/entropy"
	stm "github.com/sdaros/passgo/stamper"
)

// Env is provided as an environment by that is accessible
// to all clients that require its functionality.
type Env struct {
	Logger
	ent.Entropy
	stm.Stamper
}

// Initialise the environment.
func Environment(logger Logger, entropy ent.Entropy, stamper stm.Stamper) *Env {
	// nil logger does nothing.
	if logger == nil {
		logger = new(NullLogger)
	}
	// nil entropy defaults to CryptoRand implementation.
	if entropy == nil {
		entropy = ent.CryptoRand
	}
	// nil stamper defaults to Scrypt implementation.
	if stamper == nil {
		stamper = stm.ScryptStamper
	}

	return &Env{Logger: logger, Entropy: entropy, Stamper: stamper}
}

func Null() *Env {
	return Environment(nil, nil, nil)
}
