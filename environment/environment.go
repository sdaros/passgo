package environment

import (
	ent "github.com/sdaros/passgo/entropy"
)

// Env is provided as an environment by that is accessible
// to all clients that require its functionality.
type Env struct {
	Logger
	ent.Entropy
	*Registrar
}

// Initialise the environment.
func Environment(logger Logger, entropy ent.Entropy, registrar *Registrar) *Env {
	// TODO: Read from config file

	// nil logger does nothing
	if logger == nil {
		logger = new(NullLogger)
	}
	// nil entropy defaults to CryptoRand implementation
	if entropy == nil {
		entropy = ent.CryptoRand
	}
	// nil registrar creates and initializes the registrar's values
	if registrar == nil {
		registrar = new(Registrar)
		registrar.values = make(map[string]interface{})

	}

	return &Env{Logger: logger, Entropy: entropy, Registrar: registrar}
}

func Null() *Env {
	return Environment(nil, nil, nil)
}
