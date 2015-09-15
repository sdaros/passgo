package sealer

import (
	_ "golang.org/x/crypto/nacl/secretbox"
)

type NaclSecretbox struct {
	Options []string
}
type Foo struct {
	Options []string
}

func (naclSecretbox *NaclSecretbox) Seal() (envelope []byte) {
	// TODO: implement nacl/secretbox
	return []byte(nil)
}
