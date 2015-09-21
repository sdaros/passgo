// Package Stamper implements a password-based key derivation function to stamp user-supplied content, for example a password, and returns the resulting hash.
package stamper

import (
	"errors"
)

var  (
	ErrStamp = errors.New("stamper: stamping failed!")
	ErrUnrecognizedImplementation = errors.New("stamper: unrecognized implementation chosen")
)

// Bulla is the name for the hash returned by the stamp function.
// The Bulla can be used as a `key` in symmetric encryption.
type Bulla struct {
	salt []byte
	content []byte
}

type stamp func(...interface{}) (*Bulla, error)

// TODO: implement Params []string as second parameter
func Use(implementation interface{}) (stamp, error) {
	switch t := implementation.(type) {
	default:
		return nil, ErrUnrecognizedImplementation
	case *Scrypt:
		return t.Stamp, nil
	}
}
