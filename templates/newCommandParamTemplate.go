package main

var newFlagTemplate string = `package cmd

import (
        "fmt"
)

type {{.Name}} struct {
	name      string
	usage     string
	value     int
}

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{
		name:      "{{.HyphenName}}",
		usage:     "{{.Usage}}",
		value:     {{.DefaultValue}},
	}
}

func (self *{{.Name}}) Set(value string) (err error) {
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

func (self *{{.Name}}) Validate() (err error) {
	// TODO: implement logic
	return nil
}

func (self *{{.Name}}) String() string {
	return fmt.Sprint(self.value)
}

func (self *{{.Name}}) Name() string { return self.name }

func (self *{{.Name}}) Usage() string { return self.usage }

func (self *{{.Name}}) Value() string { return self.value }
`
