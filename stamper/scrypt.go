package stamper

import (
	sc "golang.org/x/crypto/scrypt" 
	"crypto/rand"
)
type Scrypt struct {
	Params []string
}
const (
	saltSize = 32
)
func (scrypt *Scrypt) Stamp(content []byte) (*Bulla, error) {
	// TODO: Convert to MAC using key
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	result, err := sc.Key(content, salt, 65536, 8, 1, 32)
	if err != nil {
		return nil, err
	}
	return &Bulla{salt: salt, content: result}, nil
}
func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}