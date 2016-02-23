// NEXT: refactor to pass postage to Stamp() and use goconvey
package stamper

import (
	"bytes"
	"fmt"
	"github.com/sdaros/passgo/entropy"
	"golang.org/x/crypto/scrypt"
	"testing"
)

type testVector []byte

// testVector must implement fmt.Stringer
func (tv testVector) String() string {
	return fmt.Sprintf("%v", string(tv))
}

// Stamp Method for our Scrypt implementation should returns
// the same content as golang.org/x/crypto/scrypt.
func Test_stamp_against_crypto_scrypt_key(t *testing.T) {
	var tv testVector
	tv = []byte("StampMe!")
	scryptStamper := &Scrypt{n: 16, r: 1, p: 1, len: 32,
		entropyImplementation: entropy.CryptoRand}
	stampedByScryptStamper, err := scryptStamper.Stamp(tv)
	if err != nil {
		t.Error("stamper/scrypt returned an error: ", err)
	}
	stampedByCryptoScrypt, err := scrypt.Key([]byte(tv),
		stampedByScryptStamper.Salt, 16, 1, 1, 32)
	if err != nil {
		t.Error("crypto/scrypt returned an error: ", err)
	}
	if !bytes.Equal(stampedByScryptStamper.Content, stampedByCryptoScrypt) {
		t.Errorf("Expected ScryptStamper to have the same hash as crypto/scrypt."+
			" Got %x, instead of %x", stampedByScryptStamper.Content,
			stampedByCryptoScrypt)
	}
}

// Stamp() Method for our Scrypt implementation should
// return a stamper.ErrStamp when called with bad params.
func Test_stamp_returns_error_on_bad_input_parameters(t *testing.T) {
	var tv testVector
	tv = "StampMe!"
	scryptStamper := &Scrypt{n: 17, r: 1, p: 1, len: 32,
		entropyImplementation: entropy.CryptoRand}
	_, err := scryptStamper.Stamp(tv)
	if err == nil {
		t.Error("Expected an error on bad params to " +
			"Scrypt Implementation, got nil error")
	}
}
