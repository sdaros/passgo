package sealer

import (
	_ "golang.org/x/crypto/nacl/secretbox"
)

type NaclSecretbox struct {
	Options []string
}

func (naclSecretbox *NaclSecretbox) Seal(secret []byte) (content []byte) {
	// TODO: implement nacl/secretbox
	return secret
}
