package sealer

type Envelope struct {
  Metadata map[string][]interface{}
  Content []byte
}
