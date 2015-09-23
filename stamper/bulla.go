package stamper

// Bulla is the name for the hash returned by the stamp function.
// The Bulla can be used as a `key` in symmetric encryption.
type Bulla struct {
  Content []byte
  Salt []byte
}
