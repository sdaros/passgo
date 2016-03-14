package cmd

import (
	"github.com/sdaros/passgo/app"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"unicode"
)

func TestPasswordRetrievesCommandFlagsFromPassgoRegistrar(t *testing.T) {
	Convey("Given a Password Command that receives its command flags "+
		"from the Passgo registrar", t, func() {
		passgo := app.Null()
		passgo = registerPasswordCommandFlagsWithPassgoRegistrar(passgo)
		passwordCmd := NewPassword()

		Convey("when we apply those flags to the Password Command", func() {
			err := passwordCmd.ApplyCommandParamsFrom(passgo)

			Convey("we should not receive an error", func() {

				So(err, ShouldBeNil)
			})

			Convey("the value of the password-length flag in Password should equal "+
				"what was provided to the Passgo registrar", func() {

				So(passwordCmd.passwordLength.value, ShouldEqual, 10)
			})

			Convey("the value of the no-symbols flag in Password should equal "+
				"what was provided to the Passgo registrar", func() {

				So(passwordCmd.noSymbols.value, ShouldEqual, false)
			})
		})

	})
}

func TestPasswordAppliesCommandFlagsProperly(t *testing.T) {
	Convey("Given a Password Command", t, func() {
		command := NewPassword()
		commandExecuteFunc := command.execute

		Convey("When the execute function is called", func() {
			_, err := commandExecuteFunc()

			Convey("no error should be returned", func() {

				So(err, ShouldBeNil)
			})

		})
		Convey("When command flags are not provided", func() {
			cmdResult, err := commandExecuteFunc()

			Convey("It should apply the default value for the password-length flag", func() {

				So(err, ShouldBeNil)
				So(len(cmdResult.(string)), ShouldEqual, 15)

			})
			Convey("It should apply the default value for the no-symbols flag", func() {
				passwordContainsSymbols, err := passwordContainsSymbols(command)

				So(err, ShouldBeNil)
				So(passwordContainsSymbols, ShouldEqual, true)
			})

		})
		Convey("When the value of password-length flag is 256", func() {
			command.passwordLength.value = 256

			Convey("The result of the command should be a password "+
				"with length of 256 characters", func() {
				cmdResult, err := commandExecuteFunc()

				So(err, ShouldBeNil)
				So(len(cmdResult.(string)), ShouldEqual, 256)

			})

		})

		Convey("When the value of no-symbols flag is true", func() {
			command.noSymbols.value = true

			Convey("The result of the command should be a password "+
				"that does not contain any symbols (is "+
				"only alphanumeric [A-Za-z])", func() {
				passwordContainsSymbols, err := passwordContainsSymbols(command)

				So(err, ShouldBeNil)
				So(passwordContainsSymbols, ShouldEqual, false)
			})
		})

		Convey("When the value of no-symbols flag is false", func() {
			command.noSymbols.value = false

			Convey("The result of the command should be a password "+
				"that *contains* symbols (see cmd.PasswordCommand)", func() {
				passwordContainsSymbols, err := passwordContainsSymbols(command)

				So(err, ShouldBeNil)
				So(passwordContainsSymbols, ShouldEqual, true)
			})
		})

	})

}

func passwordContainsSymbols(p *Password) (bool, error) {
	// password-length should be sufficiently large.
	p.passwordLength.value = 256
	passwordExecuteFn := p.ExecuteFn()
	cmdResult, err := passwordExecuteFn()
	if err != nil {
		return false, err
	}
	isSymbolFound := false
	for _, r := range cmdResult.(string) {
		if unicode.IsSymbol(r) {
			isSymbolFound = true
			break
		}
	}
	return isSymbolFound, nil
}

func registerPasswordCommandFlagsWithPassgoRegistrar(passgo *app.App) *app.App {
	plengthFlag := NewPasswordLengthFlag()
	plengthFlag.value = 10
	passgo.Register(plengthFlag.Name(), plengthFlag)

	noSymbolsFlag := NewNoSymbolsFlag()
	passgo.Register(noSymbolsFlag.Name(), noSymbolsFlag)

	return passgo
}
