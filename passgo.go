package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/sdaros/passgo/stamper"
)

func main() {
	facebook := &secret{"https://facebook.com", "p@ssw0rd", "username", "note"}
	content, err := facebook.Seal()
	if err != nil {
		panic(err)
	}
	result := new(secret)
	stampedResult, err := facebook.Stamp()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Secret: %v,\nSealed Secret: %v,\nStamped Secret: %v\n",
		facebook, result, stampedResult)
}
