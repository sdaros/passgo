package cmd

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"unicode"
)

// NEXT: finish test suite with goconvey
type testVector string

// testVector must implement fmt.Stringer
func (tv testVector) String() string {
	return fmt.Sprintf("%v", string(tv))
}

func TestPasswordAppliesCommandFlagsProperly(t *testing.T) {
	Convey("Given a Password Command to execute", t, func() {
		command := NewPassword()
		commandExecuteFunc := command.execute

		Convey("When no flags to the command are passed", func() {
			cmdResult, err := commandExecuteFunc()

			Convey("It should apply the default values of the flags", func() {

				So(err, ShouldBeNil)
				So(len(cmdResult.String()), ShouldEqual, 15)
				So(passwordContainsSymbols(cmdResult), ShouldEqual, true)

			})

		})
		Convey("When the value of password-length flag is 256", func() {
			command.passwordLength.value = 256

			Convey("The result of the command should be a password "+
				"with length of 256 characters", func() {
				cmdResult, err := commandExecuteFunc()

				So(err, ShouldBeNil)
				So(len(cmdResult.String()), ShouldEqual, 256)

			})

		})

		Convey("When the value of no-symbols flag is true", func() {
			command.noSymbols.value = true

			Convey("The result of the command should be a password "+
				"that does not contain any symbols (is "+
				"only alphanumeric [A-Za-z])", func() {
				cmdResult, err := commandExecuteFunc()

				So(err, ShouldBeNil)
				So(passwordContainsSymbols(cmdResult), ShouldEqual, false)
			})
		})

		Convey("When the value of no-symbols flag is false", func() {
			command.noSymbols.value = false

			Convey("The result of the command should be a password "+
				"that *contains* symbols (see cmd.PasswordCommand)", func() {
				cmdResult, err := commandExecuteFunc()

				So(err, ShouldBeNil)
				So(passwordContainsSymbols(cmdResult), ShouldEqual, true)
			})
		})

	})

}

func passwordContainsSymbols(c *CmdResult) bool {
	isSymbolFound := false
	for _, r := range c.String() {
		if unicode.IsSymbol(r) {
			isSymbolFound = true
			break
		}
	}
	return isSymbolFound
}

func Test_password_with_default_flags(t *testing.T) {
	command := NewPassword()
	exec := command.execute
	_, err := exec()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
