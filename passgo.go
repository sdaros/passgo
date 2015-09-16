package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/sdaros/passgo/stamper"
)

func main() {
	facebook := &secret{"https://facebook.com", "p@ssw0rd", "username", "note"}
	sealedSecret := facebook.Seal()
	result := new(secret)
	stampedResult := facebook.Stamp()
	err := json.Unmarshal(sealedSecret, result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Secret: %v,\nSealed Secret: %v,\nStamped Secret: %v\n",
		facebook, result, stampedResult)
}
