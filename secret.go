package passgo

import (
	"encoding/json"
	_ "golang.org/x/crypto/nacl/secretbox"
)

type secret struct {
	Url      string
	Password string
	Username string
	Note     string
}

func (sec secret) String() string {
	result, err := json.MarshalIndent(sec, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(result[:])
}

func (secret *secret) seal() []byte {
	// TODO: implement nacl/secretbox
	return []byte(nil)
}
