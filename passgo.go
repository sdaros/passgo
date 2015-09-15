package main
import (
	"encoding/hex"
	"fmt"
)

func main() {
	secret := &secret{"https://facebook.com", "p@ssw0rd", "username", "note"}
	tag := tag("facebook")
	stamp := &stamp{stamper.lick(tag)}
	fmt.Printf("Secret: %v,\nSealed Secret (Stamp): %v,\n",
		secret, hex.EncodeToString(stamp.content))
	fmt.Printf("Seal: %v", secret.Seal())
}
