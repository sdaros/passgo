package sealer

import (
	"errors"
	"fmt"
)

var (
	ErrSeal = errors.New("sealer: sealing failed!")
	ErrOpen = errors.New("sealer: opening failed!")
	ErrUnrecognizedImplementation = errors.New("sealer: unrecognized implementation chosen")
)

// TODO: should return a sealedSecret instead of a []byte
type seal func([]byte) ([]byte, error)

// TODO: implement options []string as second parameter
func Use(implementation interface{}) seal {
	switch t := implementation.(type) {
	default:
		panic(fmt.Sprintf("%v (%T)", ErrUnrecognizedImplementation, t))
	case *NaclSecretbox:
		return t.Seal
	}
}