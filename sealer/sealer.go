package sealer

import (
	"errors"
)

// TODO: return these errors eventually
var (
	ErrSeal = errors.New("sealer: sealing failed!")
	ErrOpen = errors.New("sealer: opening failed!")
)

// postage is sealed (authenticated encryption) by a sealer implementation
// postage is usually a secret represented as a JSON string
type postage interface {
	String() string
}

// naclSecretboxSealer uses nacl/secretbox for symmetric encryption
var NaclSecretboxSealer = new(NaclSecretbox)
