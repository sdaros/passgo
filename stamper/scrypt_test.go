package stamper

import (
  "golang.org/x/crypto/scrypt"
  "encoding/hex"
  "testing"
  "bytes"
  "fmt"

)
type testVector string
// testVector must implement fmt.Stringer
func (tv testVector) String() string {
    return fmt.Sprintf("%v", string(tv))
}
// Test that the Stamp() Method for our Scrypt implementation
// returns the same content as golang.org/x/crypto/scrypt.
func Test_stamp_against_crypto_scrypt_key(t *testing.T) {
  var tv testVector
  tv = "StampMe!"
  scryptStamper := &Scrypt{n: 16, r: 1, p: 1, len: 32}
  stampedByScryptStamper, err := scryptStamper.Stamp(tv)
  if err != nil {
    t.Error("stamper/scrypt returned an error: ", err)
  }
  stampedByCryptoScrypt, err := scrypt.Key([]byte(tv),
    stampedByScryptStamper.Salt, 16, 1, 1, 32)
  if err != nil {
    t.Error("crypto/scrypt returned an error: ", err)
  }
  if ! bytes.Equal(stampedByScryptStamper.Content, stampedByCryptoScrypt)  {
    t.Errorf("Expected ScryptStamper to have the same hash as crypto/scrypt." +
      " Got %x, instead of %x", stampedByScryptStamper.Content,
      stampedByCryptoScrypt)
  }
}
// Test that the Stamp() Method for our Scrypt implementation
// returns a stamper.ErrStamp when called with bad params.
func Test_stamp_returns_error_on_bad_input_parameters(t *testing.T) {

}
