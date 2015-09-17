package main

import (
	"github.com/sdaros/passgo/stamper"
	"github.com/sdaros/passgo/sealer"
	"encoding/json"
	"log"
	"fmt"
)

type secret struct {
	Url      string
	Password string
	Username string
	Note     string
}

func (secret *secret) String() (string) {
	jsonString, err := json.MarshalIndent(secret, "", "\t")
	if err != nil {
		log.Fatalln(err)
		return fmt.Sprint(secret)
	}
	return string(jsonString[:])
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
	bulla, err :=  stamp([]byte(secret.String()))
	if err != nil {
		return nil, err
	}
	return bulla, nil
}
