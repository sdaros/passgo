package cmd

import (
	"errors"

	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cmd/stamp"
	"github.com/sdaros/passgo/stamper"
)

// Stamp returns stamped (hashed with a PBKDF) postage.
type Stamp struct {
	*app.App
	execute func() (CmdResult, error)
	impl    stamper.Stamper
	name    string
	postage *stamp.Postage
	result  CmdResult
}

// NewStamp returns a stamp command with default values.
func NewStamp() *Stamp {
	stamp := &Stamp{
		App:     app.Null(),
		impl:    NewScrypt(),
		name:    "stamp",
		postage: stamp.NewPostage(),
	}
	stamp.execute = stampExecuteFn(stamp)
	return stamp
}

// stampExecuteFn returns a Bulla which is the result of stamping
// the postage plus associated salt.
func stampExecuteFn(s *Stamp) func() (CmdResult, error) {
	stampExecuteFn := func() (CmdResult, error) {
		s.ApplyCommandParamsFrom(s.App)
		if err := s.validate(); err != nil {
			return nil, err
		}
		bulla, err := s.impl.Stamp(s.postage.Value())
		if err != nil {
			return nil, err
		}
		return bulla, nil
	}
	return stampExecuteFn
}

func (s *Stamp) ExecuteFn() func() (CmdResult, error) { return s.execute }

func (s *Stamp) ApplyCommandParamsFrom(passgo *app.App) error {
	if passgo == nil {
		return errors.New("We need a valid Passgo object to retrieve flags")
	}
	s.App = passgo
	if s.App.Lookup("postage") != nil {
		postageFromApp := s.App.Lookup("postage").(*stamp.Postage)
		s.postage = postageFromApp
	} // else, postage param was not provided.
	return nil
}

func (s *Stamp) validate() error {
	return nil
}

func (s *Stamp) Name() string { return s.name }
