package stamper

import (
	"golang.org/x/crypto/scrypt"
	"crypto/rand"
)
type Scrypt struct {
	// TODO: Implement args
}

const (
	saltSize = 32
)

func (s *Scrypt) Stamp(postage postage) (*Bulla, error) {
	var content []byte
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	copy(content[:], []byte(postage.String()))
	result, err := scrypt.Key(content, salt, 65536, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	return &Bulla{Salt: salt, Content: result}, nil
}
func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
