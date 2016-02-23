package stamper

import (
	"github.com/sdaros/passgo/entropy"
	"github.com/sdaros/passgo/mailbag"
	"golang.org/x/crypto/scrypt"
)

type Scrypt struct {
	// n and r control scrypt's memory requirements
	n int
	r int
	// p controls whether scrypt can run on multiple processors
	p int
	// length in bytes
	len                   int
	entropyImplementation entropy.Entropy
}

func (s *Scrypt) Stamp(postage mailbag.Postage) (*mailbag.Bulla, error) {
	salt, err := generateSalt(s.len, s.entropyImplementation)
	if err != nil {
		return nil, ErrStamp
	}
	result, err := scrypt.Key(postage, salt, s.n, s.r, s.p, s.len)
	if err != nil {
		return nil, ErrStamp
	}
	return &mailbag.Bulla{Salt: salt, Content: result}, nil
}

func generateSalt(saltLength int, ent entropy.Entropy) ([]byte, error) {
	salt := make([]byte, saltLength)
	_, err := ent.Read(salt)
	if err != nil {
		return nil, ErrStamp
	}
	return salt, nil
}
