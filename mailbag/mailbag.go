package mailbag

type Postage []byte

func NewPostage() *Postage {
	return new(Postage)
}
