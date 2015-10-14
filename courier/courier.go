// Package Courier Proxies all requests from the client
// for example, seal() and stamp() ...
package courier

import (
	"golang.org/x/crypto/ssh/terminal"
	_"github.com/sdaros/passgo/stamper"
	_"github.com/sdaros/passgo/sealer"
	_"github.com/sdaros/passgo/cmd"
	"os"
)

func Password() ([]byte, error) {
	fd := int(os.Stdin.Fd())
	pass, err := terminal.ReadPassword(fd)
	if err != nil {
		return nil, err
	}
	return pass, nil
}
