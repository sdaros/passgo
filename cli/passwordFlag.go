package cli

type passwordFlag struct {
	name      string
	usage     string
	isCommand bool
}

func NewPasswordFlag() *passwordFlag {
	return &passwordFlag{
		name:      "password",
		usage:     "Generate a new password",
		isCommand: true,
	}
}

func (self *passwordFlag) Set(value string) (err error) { return nil }

func (self *passwordFlag) String() string { return "" }

func (self *passwordFlag) IsBoolFlag() bool { return true }

func (self *passwordFlag) IsCommand() bool { return self.isCommand }

func (self *passwordFlag) Name() string { return self.name }

func (self *passwordFlag) Usage() string { return self.usage }
