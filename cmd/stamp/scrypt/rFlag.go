package cmd

import (
	"fmt"
	"strconv"
)

type rFlag struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewRFlag() *rFlag {
	return &rFlag{
		name:      "scrypt-r-flag",
		usage:     "Controls scrypt memory use.",
		value:     8,
		isCommand: false,
	}
}

func (self *rFlag) Set(value string) (err error) {
	// TODO: implement logic
	rParam, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(); err != nil {
		return err
	}
	self.value = rParam
	return nil
}

func (self *rFlag) Validate() (err error) {
	// TODO: implement logic
	return nil
}

func (self *rFlag) String() string {
	return fmt.Sprint(self.value)
}

func (self *rFlag) IsCommand() bool { return self.isCommand }

func (self *rFlag) Name() string { return self.name }

func (self *rFlag) Usage() string { return self.usage }
