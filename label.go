package main

import (
	"encoding/json"
)

type Label struct {
	Name string
}

func (label *Label) String() string {
	jsonString, err := json.Marshal(label)
	if err != nil {
		panic(err)
	}
	return string(jsonString[:])
}
