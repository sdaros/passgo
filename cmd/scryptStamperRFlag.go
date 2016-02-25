package cmd

import (
	"fmt"
	"strconv"
)

type scryptStamperRFlag struct {
	name      string `schema.org: "/name"`
	usage     string `schema.org: "/description"`
	value     int    `schema.org: "/value"`
	isCommand bool
}

// NewScryptStamperRFlag returns a ScryptStamperR flag with a default value.
func NewScryptStamperRFlag() *scryptStamperRFlag {
	return &scryptStamperRFlag{
		name:      "scrypt-stamper-r",
		usage:     "parameter `r` to control scrypt's memory requirements.",
		value:     8,
		isCommand: false,
	}
}

func (st *scryptStamperRFlag) Name() string {
	return st.name
}

func (st *scryptStamperRFlag) Usage() string {
	return st.usage
}

func (st *scryptStamperRFlag) IsCommand() bool {
	return st.isCommand
}

// String is provided to satisfy flag.Value interface.
func (st *scryptStamperRFlag) String() string {
	return fmt.Sprint(st.value)
}

// Set sets the value for the scryptStamperRFlag and validates the range.
func (st *scryptStamperRFlag) Set(value string) (err error) {
	r, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := st.Validate(r); err != nil {
		return err
	}
	st.value = r
	return nil
}

func (st *scryptStamperRFlag) Validate(r int) (err error) {
	//TODO
	return nil
}
