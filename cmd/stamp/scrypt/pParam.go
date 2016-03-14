package scrypt

import (
	"fmt"
	"strconv"
)

type PParam struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewPParam() *PParam {
	return &PParam{
		name:      "scrypt-p-param",
		usage:     "controls if scrypt should run on multiple CPU cores.",
		value:     1,
		isCommand: false,
	}
}

func (self *PParam) Set(value string) (err error) {
	// TODO: implement logic
	pParam, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(pParam); err != nil {
		return err
	}
	self.value = pParam
	return nil
}

func (self *PParam) Validate(value interface{}) (err error) {
	return nil
}

func (self *PParam) String() string {
	return fmt.Sprint(self.value)
}

func (self *PParam) IsCommand() bool { return self.isCommand }

func (self *PParam) Name() string { return self.name }

func (self *PParam) Usage() string { return self.usage }

func (self *PParam) Value() int { return self.value }
