// Package Courier Proxies all requrest from the client 
// for example, seal() and stamp() ...

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)
package courier

// Label can be used to tag the contents of an envelope.
type label string

// Envelope can contain sealed secrets and any associated labels.
type envelope []byte

func password() []byte {
	fd := int(os.Stdin.Fd())
	pass, err := terminal.ReadPassword(fd)
	if err != nil {
		panic(err)
	}
	return pass
}
