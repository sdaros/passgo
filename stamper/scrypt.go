// Packages  Stamper implements a password-based key derivation function
// to stamp user-supplied content (ex. password) into a bulla.
package stamper

import (
	sc "golang.org/x/crypto/scrypt" 
	"golang.org/x/crypto/sha3"
)

type Scrypt struct {
	Options []string
}
func (scrypt *Scrypt) Stamp(content []byte) (Bulla) {
	// TODO: Convert to MAC using key
	salt := make([]byte, 64)
	sha3.ShakeSum256(salt, content)
	result, err := sc.Key(content, salt, 65536, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	return Bulla(result)
}