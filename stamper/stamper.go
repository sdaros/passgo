// Stamper implements a password-based key derivation function
// to stamp user-supplied content, for example to provide a
// secure hash of a user supplied password.
package stamper

import (
	"errors"
	"fmt"
	"github.com/sdaros/passgo/mailbag"
)

type Stamper interface {
	Stamp(mailbag.Postage) (*mailbag.Bulla, error)
}

var ErrStamp = errors.New("stamper: stamping failed, check your input parameters.")

// postage is hashed (using PBKDF) by a stamper implementation.
// postage is usually a Label represented as a JSON string.
type postage interface {
	fmt.Stringer
}
