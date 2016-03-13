package cmd

type generateFlag struct {
	isCommand bool
	name      string
	usage     string
}

func NewGenerateFlag() *generateFlag {
	return &generateFlag{
		isCommand: true,
		name:      "generate",
		usage:     "Generate a new password",
	}
}

func (self *generateFlag) Set(value string) (err error) { return nil }

func (self *generateFlag) String() string { return "" }

func (self *generateFlag) IsBoolFlag() bool { return true }

func (self *generateFlag) IsCommand() bool { return self.isCommand }

func (self *generateFlag) Name() string { return self.name }

func (self *generateFlag) Usage() string { return self.usage }
