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

func (secret *secret) Seal() (sealedSecret []byte) {
	implementation := new(sealer.NaclSecretbox)
	seal := sealer.Use(implementation)
	return seal([]byte(secret.String()))
}
func (secret *secret) Stamp() (stamper.Bulla) {
	implementation := new(stamper.Scrypt)
	stamp := stamper.Use(implementation)
	return stamp([]byte(secret.String()))
}
