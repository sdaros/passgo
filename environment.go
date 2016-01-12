package main

import (
	"github.com/sdaros/passgo/courier"
	ent "github.com/sdaros/passgo/entropy"
)

// Env is provided as an environment by courier that is accessible
// to all clients that require its functionality
type Env struct {
	courier.Logger
	ent.Entropy
}

// Initialise the environment
func Environment(logger courier.Logger, entropy ent.Entropy) *Env {
	// TODO: Read from config file

	// nil logger does nothing
	if logger == nil {
		logger = new(courier.NullLogger)
	}
	// nil entropy default to CryptoRand implementation
	if entropy == nil {
		entropy = ent.CryptoRand
	}
	return &Env{Logger: logger, Entropy: entropy}
}
