package stamp

import (
	"fmt"
	"github.com/sdaros/passgo/mailbag"
)

type Postage struct {
	name      string
	usage     string
	value     mailbag.Postage
	isCommand bool
}

func NewPostage() *Postage {
	return &Postage{
		name:      "postage",
		usage:     "postage that will be stamped.",
		isCommand: false,
	}
}

func (self *Postage) Set(value string) (err error) {
	if err := self.Validate(nil); err != nil {
		return err
	}
	self.value = mailbag.Postage([]byte(value))
	return nil
}

func (self *Postage) Validate(value interface{}) (err error) {
	return nil
}

func (self *Postage) String() string {
	return fmt.Sprint(self.value)
}

func (self *Postage) IsCommand() bool { return self.isCommand }

func (self *Postage) Name() string { return self.name }

func (self *Postage) Usage() string { return self.usage }

func (self *Postage) Value() mailbag.Postage { return self.value }
