package main
import (
	"encoding/json"
)

type Secret struct {
	Url      string
	Password string
	Username string
	Note     string
}

func (secret *Secret) String() string {
	jsonString, err := json.Marshal(secret)
	if err != nil {
		panic(err)
	}
	return string(jsonString[:])
}
