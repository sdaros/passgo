package sealer

import (
	_ "golang.org/x/crypto/nacl/secretbox"
)

type NaclSecretbox struct {
	// TODO: implement args
}

func (secretbox *NaclSecretbox) Seal(post postage) (env *Envelope, err error) {
	// TODO: implement nacl/secretbox
	return &Envelope{}, nil
}
