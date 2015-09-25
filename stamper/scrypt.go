package stamper

import (
	"golang.org/x/crypto/scrypt"
	"crypto/rand"
)
type Scrypt struct {
	// n and r control scrypt's memory requirements
	n int
	r int
	// p controls whether scrypt can run on multiple processors
	p int
	// length in bytes
	len int
}

func (s *Scrypt) Stamp(postage postage) (*Bulla, error) {
	salt, err := generateSalt(s.len)
	if err != nil {
		return nil, err
	}
	result, err := scrypt.Key([]byte(postage.String()), salt, s.n, s.r, s.p, s.len)
	if err != nil {
		return nil, err
	}
	return &Bulla{Salt: salt, Content: result}, nil
}

func generateSalt(saltLength int) ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
