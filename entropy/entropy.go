// entropy provides basic functions to retrieve data from
// an entropy pool and map it into a character set
package entropy

import (
	"io"
)

type Entropy interface {
	// adds random bytes from the entropy pool.
	io.Reader
	// selects a random integer from a defined range.
	Int(max int) (n int64, err error)
}

// CryptoRand uses crypto/rand in the standard library
// to collect entropy
var CryptoRand = new(cryptoRand)
