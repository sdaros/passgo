package cmd

import (
	"errors"

	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cmd/stamp/scrypt"
	"github.com/sdaros/passgo/mailbag"
	"github.com/sdaros/passgo/stamper"
)

type Scrypt struct {
	*app.App
	execute func() (CmdResult, error)
	length  *scrypt.Length
	name    string
	nParam  *scrypt.NParam
	pParam  *scrypt.PParam
	rParam  *scrypt.RParam
}

func NewScrypt() *Scrypt {
	scrypt := &Scrypt{
		App:    app.Null(),
		length: scrypt.NewLength(),
		name:   "scrypt",
		nParam: scrypt.NewNParam(),
		pParam: scrypt.NewPParam(),
		rParam: scrypt.NewRParam(),
	}
	scrypt.execute = scryptExecuteFn(scrypt)
	return scrypt
}

func scryptExecuteFn(sc *Scrypt) func() (CmdResult, error) {
	scryptExecuteFn := func() (CmdResult, error) {
		sc.ApplyCommandParamsFrom(sc.App)
		if err := sc.Validate(nil); err != nil {
			return nil, err
		}
		return sc.Stamp(new(mailbag.Postage))
	}
	return scryptExecuteFn
}

func (sc *Scrypt) Stamp(postage *mailbag.Postage) (*mailbag.Bulla, error) {
	stamper := &stamper.Scrypt{
		EntropyImplementation: sc.App.Entropy,
		Length:                sc.length.Value(),
		N:                     sc.nParam.Value(),
		P:                     sc.pParam.Value(),
		R:                     sc.rParam.Value(),
	}
	bulla, err := stamper.Stamp(postage)
	if err != nil {
		return nil, err
	}
	return bulla, nil
}

func (sc *Scrypt) ExecuteFn() func() (CmdResult, error) { return sc.execute }

func (sc *Scrypt) ApplyCommandParamsFrom(passgo *app.App) error {
	if passgo == nil {
		return errors.New("We need a valid Passgo object to retrieve flags")
	}
	sc.App = passgo
	if sc.App.Lookup("scrypt-n-param") != nil {
		nParamFromApp := sc.App.Lookup("scrypt-n-param").(*scrypt.NParam)
		sc.nParam = nParamFromApp
	} // else, scrypt-n-param was not provided; so the default will be used.
	if sc.App.Lookup("scrypt-r-param") != nil {
		rParamFromApp := sc.App.Lookup("scrypt-r-param").(*scrypt.RParam)
		sc.rParam = rParamFromApp
	} // else, scrypt-r-param was not provided; so the default will be used.
	if sc.App.Lookup("scrypt-p-param") != nil {
		pParamFromApp := sc.App.Lookup("scrypt-p-param").(*scrypt.PParam)
		sc.pParam = pParamFromApp
	} // else, scrypt-p-param was not provided; so the default will be used.
	if sc.App.Lookup("scrypt-length") != nil {
		lengthFromApp := sc.App.Lookup("scrypt-length").(*scrypt.Length)
		sc.length = lengthFromApp
	} // else, scrypt-length was not provided; so the default will be used.
	return nil
}

func (self *Scrypt) Validate(value interface{}) (err error) {
	return nil
}

func (self *Scrypt) Name() string { return self.name }
