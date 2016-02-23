package mailbag

type (
	// Postage is stampable.
	Postage []byte
	// Bulla is the name for the hash returned by the stamp function.
	// The Bulla can be used as a `key` in symmetric encryption.
	Bulla struct {
		Content []byte
		Salt    []byte
	}
)
