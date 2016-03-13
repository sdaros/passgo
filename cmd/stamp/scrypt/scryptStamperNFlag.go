package cmd

import (
        "fmt"
	"strconv"
)

type scryptStamperNFlag struct {
	name      string
	usage     string
	value     int
	isCommand bool
}

func NewScryptStamperNFlag() *scryptStamperNFlag {
	return &scryptStamperNFlag{
		name:      "scrypt-n-param",
		usage:     "controls the scrypt n parameter.",
		value:     65536,
		isCommand: false,
	}
}

func (self *scryptStamperNFlag) Set(value string) (err error) {
	// TODO: implement logic
	nFlag, err := strconv.Atoi(value)
	if err != nil {
		return err
	if err := self.Validate(); err != nil {
		return err
	}
	self.value = nFlag
	return nil
}

func (self *scryptStamperNFlag) Validate() (err error) {
	// TODO: implement logic
	return nil
}

func (self *scryptStamperNFlag) String() string {
	return fmt.Sprint(self.value)
}

func (self *scryptStamperNFlag) IsCommand() bool { return self.isCommand }

func (self *scryptStamperNFlag) Name() string { return self.name }

func (self *scryptStamperNFlag) Usage() string { return self.usage }
