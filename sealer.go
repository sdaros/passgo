package main

import (
	_ "golang.org/x/crypto/nacl/secretbox"
)

type sealer interface {
	seal() []byte
}

func (secret *secret) seal() []byte {
	// TODO: implement nacl/secretbox
	return []byte(nil)
}
