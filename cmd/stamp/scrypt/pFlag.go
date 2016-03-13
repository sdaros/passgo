package cmd

import (
	"fmt"
	"strconv"
)

type pFlag struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewPFlag() *pFlag {
	return &pFlag{
		name:      "scrypt-p-flag",
		usage:     "controls if scrypt should run on multiple CPU cores.",
		value:     1,
		isCommand: false,
	}
}

func (self *pFlag) Set(value string) (err error) {
	// TODO: implement logic
	pParam, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(); err != nil {
		return err
	}
	self.value = pParam
	return nil
}

func (self *pFlag) Validate() (err error) {
	// TODO: implement logic
	return nil
}

func (self *pFlag) String() string {
	return fmt.Sprint(self.value)
}

func (self *pFlag) IsCommand() bool { return self.isCommand }

func (self *pFlag) Name() string { return self.name }

func (self *pFlag) Usage() string { return self.usage }
