package sealer

type Envelope struct {
  Metadata map[string][]interface{}
  Message []byte
  Nonce []byte
}
