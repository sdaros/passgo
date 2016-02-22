package sealer

import (
	"errors"
	"fmt"
	"github.com/sdaros/passgo/entropy"
)

var (
	ErrSeal = errors.New("sealer: sealing failed!")
	ErrOpen = errors.New("sealer: opening failed!")
)

type Sealer interface {
	Seal(postage) (env *Envelope, err error)
}

// naclSecretboxSealer uses nacl/secretbox for symmetric encryption.
// The crypto/rand library is used as the default entropy source.
var NaclSecretboxSealer = &NaclSecretbox{entropyImplementation: entropy.CryptoRand}

// postage is sealed (encrypted then authenticated) by a sealer implementation.
// postage is usually a secret represented as a JSON string.
type postage interface {
	fmt.Stringer
}
