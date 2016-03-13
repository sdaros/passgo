package cmd

import (
        "fmt"
	"strconv"
)

type scryptStamperLenFlag struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewScryptStamperLenFlag() *scryptStamperLenFlag {
	return &scryptStamperLenFlag{
		name:      "scrypt-stamper-len-flag",
		// Insert command flags here
		usage:     "controls the length of scrypt's output (in bytes).",
		value:     32,
		isCommand: false,
	}
}

func (self *scryptStamperLenFlag) Set(value string) (err error) {
	// TODO: implement logic
	lenFlag, err := strconv.Atoi(value)
	if err != nil {
		return err
	if err := self.Validate(); err != nil {
		return err
	}
	self.value = lenFlag
	return nil
}

func (self *scryptStamperLenFlag) Validate() (err error) {
	// TODO: implement logic
	return nil
}

func (self *scryptStamperLenFlag) String() string {
	return fmt.Sprint(self.value)
}

func (self *scryptStamperLenFlag) IsCommand() bool { return self.isCommand }

func (self *scryptStamperLenFlag) Name() string { return self.name }

func (self *scryptStamperLenFlag) Usage() string { return self.usage }
