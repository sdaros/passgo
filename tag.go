package passgo

import (
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/sha3"
)

type tag string

func (tag tag) lick() []byte {
	// TODO: Convert to MAC using key
	salt := make([]byte, 64)
	sha3.ShakeSum256(salt, []byte(tag))
	content, err := scrypt.Key([]byte(tag), salt, 65536, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return content
}

