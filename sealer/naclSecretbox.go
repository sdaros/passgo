package sealer

import (
	_ "golang.org/x/crypto/nacl/secretbox"
)

type NaclSecretbox struct {
	Options	[]string
}

func (naclSecretbox *NaclSecretbox) Seal() ([]byte, error) {
	// TODO: implement nacl/secretbox
	return []byte(nil), nil
}