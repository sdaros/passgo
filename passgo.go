package main

import (
	"github.com/sdaros/passgo/courier"
	_ "github.com/sdaros/passgo/entropy"
	"github.com/sdaros/passgo/sealer"
	"github.com/sdaros/passgo/stamper"
)

func main() {
	env := Environment(new(courier.StandardLogger), nil)
	if err := courier.ParseOptions(); err != nil {
		panic(err)
	}
	lbl := &Label{"https://lbl.com"}
	content, err := stamper.ScryptStamper.Stamp(lbl)
	if err != nil {
		panic(err)
	}
	env.Infof("Label: %v,\nStamped Label: %v\n", lbl, content)
	sct := &Secret{"https://facebook.com", "p@ssw0rd", "user", "foob"}
	envelope, err := sealer.NaclSecretboxSealer.Seal(sct)
	if err != nil {
		panic(err)
	}

	env.Infof("Secret: %v,\nSealed Secret: %v\n", sct, envelope)
	unsealedSecret, err := sealer.NaclSecretboxSealer.Open(envelope)
	if err != nil {
		panic(err)
	}
	env.Infof("UnsealedSecret: %v", unsealedSecret)

}
