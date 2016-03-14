package scrypt

import (
	"fmt"
	"strconv"
)

type RParam struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewRParam() *RParam {
	return &RParam{
		name:      "scrypt-r-param",
		usage:     "Controls scrypt memory use.",
		value:     8,
		isCommand: false,
	}
}
func (self *RParam) Set(value string) (err error) {
	// TODO: implement logic
	rParam, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(rParam); err != nil {
		return err
	}
	self.value = rParam
	return nil
}

func (self *RParam) Validate(value interface{}) (err error) {
	return nil
}

func (self *RParam) String() string {
	return fmt.Sprint(self.value)
}

func (self *RParam) IsCommand() bool { return self.isCommand }

func (self *RParam) Name() string { return self.name }

func (self *RParam) Usage() string { return self.usage }

func (self *RParam) Value() int { return self.value }
