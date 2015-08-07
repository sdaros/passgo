package main

import (
	"fmt"
)

type secret struct {
	url      string
	password string
	username string
	note     string
}

type envelope struct {
	content []byte
}

type stamp struct {
	content [32]byte
}

func main() {
	secret := &secret{"https://facebook.com", "p@ssw0rd", "username", "note"}
	envelope := &envelope{sealer.seal(secret)}
	fmt.Printf("Secret: %v,\n Sealed: %v,\n",
		secret.password, envelope.content)
}
