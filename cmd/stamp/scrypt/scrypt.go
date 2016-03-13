package cmd

import (
	"errors"
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/mailbag"
)

type scrypt struct {
	name                 string
	execute              func() (CmdResult, error)
	scryptStamperLenFlag *scryptStamperLenFlag
	scryptStamperNFlag   *scryptStamperNFlag
	scryptStamperPFlag   *scryptStamperPFlag
	scryptStamperRFlag   *scryptStamperRFlag
	*app.App
}

func NewScrypt() *scrypt {
	scrypt := &scrypt{
		name:      "scrypt",
		lenFlag:   NewLenFlag(),
		nFlag:     NewNFlag(),
		PFlag:     NewPFlag(),
		RFlag:     NewRFlag(),
		usage:     "generate stamped (hashed with PBKDF) postage.",
		isCommand: false,
	}
	scrypt.execute = scryptExecuteFn(scrypt)
	return scrypt
}

//
func scryptExecuteFn(sc *scrypt) func() (CmdResult, error) {
	scryptExecuteFn := func() (CmdResult, error) {
		sc.ApplyCommandFlagsFrom(sc.App)
		if err := sc.Validate(); err != nil {
			return nil, err
		}
		// TODO: remaining steps for execution here
	}
	return scryptExecuteFn
}

//
func (sc *scrypt) ExecuteFn() func() (CmdResult, error) { return sc.execute }

//
func Stamp(*mailbag.Postage) (*mailbag.Bulla, error) {
	// TODO: implementation
}

//
func (sc *scrypt) ApplyCommandFlagsFrom(passgo *app.App) error {
	if passgo == nil {
		return errors.New("We need a valid Passgo object to retrieve flags")
	}
	sc.App = passgo
	// TODO: retrieve remaining flags from passgo App
	return nil
}

//
func (self *scrypt) Validate() (err error) {
	// TODO: implement logic
	return nil
}

//
func (self *scrypt) Name() string {
	return self.name
}
