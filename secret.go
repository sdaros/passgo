package main

import (
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
