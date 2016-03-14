package scrypt

import (
	"fmt"
	"strconv"
)

type NParam struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewNParam() *NParam {
	return &NParam{
		name:      "scrypt-n-param",
		usage:     "controls the scrypt n parameter.",
		value:     65536,
		isCommand: false,
	}
}

func (self *NParam) Set(value string) (err error) {
	nFlag, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(nFlag); err != nil {
		return err
	}
	self.value = nFlag
	return nil
}

func (self *NParam) Validate(value interface{}) (err error) {
	return nil
}

func (self *NParam) String() string {
	return fmt.Sprint(self.value)
}

func (self *NParam) IsCommand() bool { return self.isCommand }

func (self *NParam) Name() string { return self.name }

func (self *NParam) Usage() string { return self.usage }

func (self *NParam) Value() int { return self.value }
