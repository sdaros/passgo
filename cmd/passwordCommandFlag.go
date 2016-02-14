package cmd

type passwordCommandFlag struct {
	name      string `schema.org: "/name"`
	usage     string
	isCommand bool
}

// NewPasswordCommandFlag returns a passwordCommandFlag with default values.
func NewPasswordCommandFlag() *passwordCommandFlag {
	return &passwordCommandFlag{
		name:      "password",
		usage:     "Generate a random password.",
		isCommand: true,
	}
}

func (p *passwordCommandFlag) Name() string {
	return p.name
}

func (p *passwordCommandFlag) Usage() string {
	return p.usage
}

func (p *passwordCommandFlag) IsCommand() bool {
	return p.isCommand
}

func (p *passwordCommandFlag) String() string {
	return ""
}

func (p *passwordCommandFlag) Set(value string) (err error) {
	return nil
}

func (p *passwordCommandFlag) IsBoolFlag() bool { return true }
