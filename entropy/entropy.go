// entropy provides basic functions to retrieve data from
// an entropy pool and map it into a character set
package entropy

import (
  "io"
)

type Entropy interface {
  // adds random bytes from the entropy pool to a byte slice.
  io.Reader
  // a randomly generated password from the entropy pool.
  Password(args ...interface{}) ([]rune, error)
}
// CryptoRand uses crypto/rand in the standard library
// to collect entropy
var CryptoRand = new(cryptoRand)
