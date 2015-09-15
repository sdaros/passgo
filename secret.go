package passgo

import (
	"github.com/sdaros/passgo/sealer"
	"encoding/json"
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
	implementation := sealer.NaclSecretbox{};
	envelope, err := implementation.Seal();
	if err != nil {
		panic(err)
	}
	return envelope
}
