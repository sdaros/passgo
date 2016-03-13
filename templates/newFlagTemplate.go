package main

var newFlagTemplate string = `package cmd

import (
        "fmt"
	"strconv"
)

type {{.LcName}} struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func New{{.Name}}() *{{.LcName}} {
	return &{{.LcName}}{
		name:      "{{.HyphenName}}",
		usage:     "{{.Usage}}",
		value:     {{.DefaultValue}},
		isCommand: false,
	}
}

func (self *{{.LcName}}) Set(value string) (err error) {
	// TODO: implement logic
	{{.ShortName}}, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(); err != nil {
		return err
	}
	self.value = {{.ShortName}}
	return nil
}

func (self *{{.LcName}}) Validate() (err error) {
	// TODO: implement logic
	return nil
}

func (self *{{.LcName}}) String() string {
	return fmt.Sprint(self.value)
}

func (self *{{.LcName}}) IsCommand() bool { return self.isCommand }

func (self *{{.LcName}}) Name() string { return self.name }

func (self *{{.LcName}}) Usage() string { return self.usage }
`
