package cmd

import (
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/mailbag"
)

// Stamp returns stamped (hashed with a PBKDF) postage.
type Stamp struct {
	name    string
	execute func() (CmdResult, error)
	postage *mailbag.Postage
	result  CmdResult
	*app.App
}

// NewStamp returns a stamp command with default values.
func NewStamp() *Stamp {
	stamp := &Stamp{
		name:    "stamp",
		postage: new(mailbag.Postage),
		App:     app.Null(),
	}
	stamp.execute = stampExecuteFn(stamp)
	return stamp
}

// stampExecuteFn returns a Bulla which is the result of stamping
// the postage plus associated salt.
func stampExecuteFn(s *Stamp) func() (CmdResult, error) {
	stampExecuteFn := func() (CmdResult, error) {
		s.ApplyCommandFlagsFrom(s.App)
		if err := s.validate(); err != nil {
			return nil, err
		}
		bulla, err := s.App.Stamp(s.postage)
		return nil, nil
	}
	return stampExecuteFn
}
