package cli

type stampFlag struct {
	name      string
	usage     string
	isCommand bool
}

func NewStampFlag() *stampFlag {
	return &stampFlag{
		name:      "stamp",
		usage:     "Stamp (hashed with PBDKF) input data",
		isCommand: true,
	}
}

func (self *stampFlag) Set(value string) (err error) { return nil }

func (self *stampFlag) String() string { return "" }

func (self *stampFlag) IsBoolFlag() bool { return true }

func (self *stampFlag) IsCommand() bool { return self.isCommand }

func (self *stampFlag) Name() string { return self.name }

func (self *stampFlag) Usage() string { return self.usage }
