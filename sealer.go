package main

import "golang.org/x/crypto/scrypt"

type sealer interface {
	seal() []byte
}

func (secret *secret) seal() []byte {
	content, err := scrypt.Key([]byte(secret.password), []byte("foo"), 1048576, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return content
}
