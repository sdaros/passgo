package passgo

import (
	"fmt"
	"encoding/hex"
)

func main() {
	secret := &secret{"https://facebook.com", "p@ssw0rd", "username", "note"}
	tag := tag("facebook")
	stamp := &stamp{stamper.lick(tag)}
	fmt.Printf("Secret: %v,\nSealed Secret (Stamp): %v,\n",
	secret, hex.EncodeToString(stamp.content))
	seal := secret.seal()
	fmt.Printf("Seal: %v", seal)
}
