package scrypt

import (
	"fmt"
	"strconv"
)

const (
	lengthMin = 32
	lengthMax = 64
)

type Length struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewLength() *Length {
	return &Length{
		name:      "scrypt-length",
		usage:     "controls the length of scrypt's output (in bytes).",
		value:     32,
		isCommand: false,
	}
}

func (self *Length) Set(value string) (err error) {
	length, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(length); err != nil {
		return err
	}
	self.value = length
	return nil
}

func (self *Length) Validate(value interface{}) (err error) {
	return nil
}

func (self *Length) String() string {
	return fmt.Sprint(self.value)
}

func (self *Length) IsCommand() bool { return self.isCommand }

func (self *Length) Name() string { return self.name }

func (self *Length) Usage() string { return self.usage }

func (self *Length) Value() int { return self.value }
