package main

var newCommandTemplate string = `package cmd

import (
        "errors"
        "github.com/sdaros/passgo/app"
)

type {{.Name}} struct {
	name      string
	execute   func() (CmdResult, error)
	// Insert command flags here
	*app.App
}

func New{{.Name}}() *{{.Name}} {
	{{.LcName}} := &{{.Name}}{
		name:      "{{.HyphenName}}",
		// Insert command flags here
		usage:     "{{.Usage}}",
		value:     {{.DefaultValue}},
		isCommand: false,
	}
	{{.LcName}}.execute = {{.LcName}}ExecuteFn({{.LcName}})
	return {{.LcName}}
}

//
func {{.LcName}}ExecuteFn({{.ShortName}} *{{.Name}}) func() (CmdResult, error) {
	{{.LcName}}ExecuteFn := func() (CmdResult, error) {
		{{.ShortName}}.ApplyCommandParamsFrom({{.ShortName}}.App)
		if err := {{.ShortName}}.Validate(); err != nil {
			return nil, err
		}
		// TODO: remaining steps for execution here
	}
	return {{.LcName}}ExecuteFn
}

//
func ({{.ShortName}} *{{.Name}}) ExecuteFn() func() (CmdResult, error) { return {{.ShortName}}.execute }

//
func ({{.ShortName}} *{{.Name}}) ApplyCommandParamsFrom(passgo *app.App) error {
	if passgo == nil {
		return errors.New("We need a valid Passgo object to retrieve flags")
	}
	// TODO: retrieve remaining flags from passgo App
	return nil
}

//
func (self *{{.Name}}) Validate() (err error) {
	// TODO: implement logic
	return nil
}

//
func (self *{{.Name}}) Name() string {
	return self.name
}
`
