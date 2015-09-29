package sealer

import (
	"errors"
	"fmt"
)

// TODO: return these errors eventually
var (
	ErrSeal = errors.New("sealer: sealing failed!")
	ErrOpen = errors.New("sealer: opening failed!")
)
// naclSecretboxSealer uses nacl/secretbox for symmetric encryption
var NaclSecretboxSealer = new(NaclSecretbox)

// postage is sealed (encrypted then authenticated) by a sealer implementation.
// postage is usually a secret represented as a JSON string.
type postage interface {
	fmt.Stringer
}
