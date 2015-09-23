package main

import (
	"fmt"
	"github.com/sdaros/passgo/stamper"
	"github.com/sdaros/passgo/sealer"
)

func main() {
	lbl := &Label{"https://lbl.com"}
	content, err := stamper.ScryptStamper.Stamp(lbl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Label: %v,\nStamped Label: %v\n", lbl, content)
	sct := &Secret{"https://facebook.com", "p@ssw0rd", "user", "foob"}
	envelope, err := sealer.NaclSecretboxSealer.Seal(sct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Secret: %v,\nSealed Secret: %v\n", sct, envelope)

}
