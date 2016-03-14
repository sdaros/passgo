package cmd

import (
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cmd/password"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"unicode"
)

func TestPasswordRetrievesCommandParamsFromPassgoRegistrar(t *testing.T) {
	Convey("Given a Password Command that receives its command params "+
		"from the Passgo registrar", t, func() {
		passgo := app.Null()
		passgo = registerPasswordCommandParamsWithPassgoRegistrar(passgo)
		command := NewPassword()
		Convey("when we apply those flags to the Password Command", func() {
			err := command.ApplyCommandParamsFrom(passgo)

			Convey("we should not receive an error", func() {

				So(err, ShouldBeNil)
			})

			Convey("the value of the password-length flag in Password should equal "+
				"what was provided to the Passgo registrar", func() {

				So(command.passwordLength.Value(), ShouldEqual, 10)
			})

			Convey("the value of the no-symbols flag in Password should equal "+
				"what was provided to the Passgo registrar", func() {

				So(command.noSymbols.Value(), ShouldEqual, false)
			})
		})

	})
}

func TestPasswordAppliesCommandParamsProperly(t *testing.T) {
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
			command.passwordLength.Set("256")

			Convey("The result of the command should be a password "+
				"with length of 256 characters", func() {
				cmdResult, err := commandExecuteFunc()

				So(err, ShouldBeNil)
				So(len(cmdResult.(string)), ShouldEqual, 256)

			})

		})

		Convey("When the value of no-symbols flag is true", func() {
			command.noSymbols.Set("true")

			Convey("The result of the command should be a password "+
				"that does not contain any symbols (is "+
				"only alphanumeric [A-Za-z])", func() {
				passwordContainsSymbols, err := passwordContainsSymbols(command)

				So(err, ShouldBeNil)
				So(passwordContainsSymbols, ShouldEqual, false)
			})
		})

		Convey("When the value of no-symbols flag is false", func() {
			command.noSymbols.Set("false")

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
	p.passwordLength.Set("256")
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

func registerPasswordCommandParamsWithPassgoRegistrar(passgo *app.App) *app.App {
	pLength := password.NewLength()
	pLength.Set("10")
	passgo.Register(pLength.Name(), pLength)

	noSymbols := password.NewNoSymbols()
	passgo.Register(noSymbols.Name(), noSymbols)

	return passgo
}
