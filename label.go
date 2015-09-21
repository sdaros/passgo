package main

import (
	"github.com/sdaros/passgo/stamper"
	"encoding/json"
	"log"
	"fmt"
)

type label struct {
	Name	string
}

func (label *label) String() (string) {
	jsonString, err := json.MarshalIndent(label, "", "\t")
	if err != nil {
		log.Fatalln(err)
		return fmt.Sprint(label)
	}
	return string(jsonString[:])
}

func (label *label) Stamp() (*stamper.Bulla, error) {
	implementation := new(stamper.Scrypt)
	stamp, err := stamper.Use(implementation)
	if err != nil {
		return nil, err
	}
	bulla, err := stamp(label)
	if err != nil {
		return nil, err
	}
	return bulla, nil
}