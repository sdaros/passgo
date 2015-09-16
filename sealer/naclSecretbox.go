package sealer

import (
	_ "golang.org/x/crypto/nacl/secretbox"
)

type NaclSecretbox struct {
	Params []string
}

func (naclSecretbox *NaclSecretbox) Seal(secret []byte) (content []byte, err error) {
	// TODO: implement nacl/secretbox
	return secret, nil
}
