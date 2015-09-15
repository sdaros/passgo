package main

import (
	"encoding/json"
	"github.com/sdaros/passgo/sealer"
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

func (secret *secret) Seal() (envelope []byte) {
	implementation := &sealer.NaclSecretbox{}
	seal := sealer.Use(implementation)
	return seal()
}
