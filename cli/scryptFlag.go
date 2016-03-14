package cli

type scryptFlag struct {
	name      string
	usage     string
	isCommand bool
}

func NewScryptFlag() *scryptFlag {
	return &scryptFlag{
		name:      "scrypt-stamper",
		usage:     "generate stamped (hashed with PBKDF) postage using scrypt.",
		isCommand: true,
	}
}

func (self *scryptFlag) Set(value string) (err error) { return nil }

func (self *scryptFlag) String() string { return "" }

func (self *scryptFlag) IsBoolFlag() bool { return true }

func (self *scryptFlag) IsCommand() bool { return self.isCommand }

func (self *scryptFlag) Name() string { return self.name }

func (self *scryptFlag) Usage() string { return self.usage }
