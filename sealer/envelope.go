package sealer

import "github.com/sdaros/passgo/stamper"

type Envelope struct {
  Metadata map[string][]interface{}
  Bulla *stamper.Bulla
  Content []byte
}
