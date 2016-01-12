// Package courier Proxies all requests from the client
// for example, seal() and stamp() ...
package courier

import (
	_ "github.com/sdaros/passgo/cmd"
	_ "github.com/sdaros/passgo/sealer"
	_ "github.com/sdaros/passgo/stamper"
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
