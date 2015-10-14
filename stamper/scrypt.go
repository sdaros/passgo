package stamper

import (
	"github.com/sdaros/passgo/entropy"
	"golang.org/x/crypto/scrypt"
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
		return nil, ErrStamp
	}
	result, err := scrypt.Key([]byte(postage.String()), salt, s.n, s.r, s.p, s.len)
	if err != nil {
		return nil, ErrStamp
	}
	return &Bulla{Salt: salt, Content: result}, nil
}

func generateSalt(saltLength int) ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := entropy.Read(salt)
	if err != nil {
		return nil, ErrStamp
	}
	return salt, nil
}
