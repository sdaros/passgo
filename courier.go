package passgo

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func getPassword() []byte {
	fd := int(os.Stdin.Fd())
	pass, err := terminal.ReadPassword(fd)
	if err != nil {
		panic(err)
	}
	return pass
}
