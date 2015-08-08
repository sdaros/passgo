package main

import (
	"fmt"
	"encoding/hex"
)

type tag string

type envelope struct {
	content []byte
}

type stamp struct {
	content []byte
}

func main() {
	secret := &secret{"https://facebook.com", "p@ssw0rd", "username", "note"}
	tag := tag("facebook")
	stamp := &stamp{stamper.lick(tag)}
	fmt.Printf("Secret: %v,\nSealed Secret (Stamp): %v,\n",
	secret, hex.EncodeToString(stamp.content))
}
