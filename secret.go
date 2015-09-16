package main

import (
	"encoding/json"
	"github.com/sdaros/passgo/sealer"
	"github.com/sdaros/passgo/stamper"
)

type secret struct {
	Url      string
	Password string
	Username string
	Note     string
}

func (secret *secret) String() string {
	result, err := json.MarshalIndent(secret, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(result[:])
}

func (secret *secret) Seal() (sealedSecret []byte, err error) {
	implementation := new(sealer.NaclSecretbox)
	seal := sealer.Use(implementation)
	sealedSecret, err = seal([]byte(secret.String()))
	if err != nil {
		return nil, err
	}
	return sealedSecret, nil
}
func (secret *secret) Stamp() (*stamper.Bulla, error) {
	implementation := new(stamper.Scrypt)
	stamp := stamper.Use(implementation)
	result, err := stamp([]byte(secret.String()))
	if err != nil {
		return nil, err
	}
	return result, nil
}
