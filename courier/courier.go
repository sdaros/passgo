// Package courier handles requests form the client to the server

package courier

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func readFromStdIn() ([]byte, error) {
	fd := int(os.Stdin.Fd())
	pass, err := terminal.ReadPassword(fd)
	if err != nil {
		return nil, err
	}
	return pass, nil
}
