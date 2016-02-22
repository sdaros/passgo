package cmd

import (
	"github.com/sdaros/passgo/stamper"
)

// Stamp returns stamped (encrypted then authenticated) postage.
type Stamp struct {
	name    string
	execute func() (*CmdResult, error)
	postage *mailbag.Postage
	result  *CmdResult
	*app.App
}

// NewStamp returns a stamp command with default values.
func NewStamp() *Stamp {
	stamp := &Stamp{
		name:    "stamp",
		postage: mailbag.NewPostage(),
		App:     app.Null(),
	}
	stamp.execute = stampExecuteFn(stamp)
}

// stampExecuteFn returns a Bulla which is the result of stamping
// the postage plus associated salt.
func stampExecuteFn(s *Stamp) func() (*CmdResult, error) {
	stampExecuteFn := func() (*CmdResult, error) {

	}

}
