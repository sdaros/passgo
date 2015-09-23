// Package Stamper implements a password-based key derivation function to stamp user-supplied content, for example a password, and returns the resulting hash.
package stamper

import (
	"errors"
)

// TODO: return this error eventually
var  ErrStamp = errors.New("stamper: stamping failed!")
// Scryptstamper uses crypto/scrypt as its PBKDF
var ScryptStamper = new(Scrypt)

// postage is hashed (using PBKDF) by a stamper implementation.
// postage is usually a Label represented as a JSON string.
type postage interface {
	String() string
}
