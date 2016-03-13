package main

var newCommandAsFlagTemplate string = `package cmd

type {{.LcName}} struct {
	name      string
	usage     string
	isCommand bool
}

func New{{.Name}}() *{{.LcName}} {
	return &{{.LcName}}{
		name:      "{{.HyphenName}}",
		usage:     "{{.Usage}}",
		isCommand: true,
	}
}

func (self *{{.Name}}) Set(value string) (err error) { return nil }

func (self *{{.Name}}) String() string { return "" }

func (self *{{.Name}}) IsBoolFlag() bool { return true }

func (self *{{.Name}}) IsCommand() bool { return self.isCommand }

func (self *{{.Name}}) Name() string { return self.name }

func (self *{{.Name}}) Usage() string { return self.usage }
`
