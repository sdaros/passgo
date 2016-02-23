package cmd

import (
	"github.com/sdaros/passgo/app"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerateCommandAppliesCommandFlagsProperly(t *testing.T) {
	Convey("Given a Generate Command to execute", t, func() {

		Convey("When user provides command flags", func() {

			Convey("Generate should pass the password-length flag and "+
				"no-symbols flag to the Password subcommand", nil)
		})
	})
}
func TestGenerateCommandRetrievesCommandFlagsFromPassgoRegistrar(t *testing.T) {
	Convey("Given a Generate Command that receives its command flags "+
		"from a Passgo registrar", t, func() {
		passgo := app.Null()
		passgo = registerGenerateCommandFlagsWithPassgoRegistrar(passgo)
		generateCmd := NewGenerate()

		Convey("when the Passgo registrar that its trying to receive "+
			"its command flags from is nil", func() {
			err := generateCmd.ApplyCommandFlagsFrom(nil)

			Convey("we should receive an error", func() {

				So(err, ShouldNotBeNil)
			})

		})

		// TODO: this convey could be compacted into a for loop
		Convey("when we apply those flags to the Generate Command", func() {
			err := generateCmd.ApplyCommandFlagsFrom(passgo)

			Convey("we should not receive an error", func() {

				So(err, ShouldBeNil)
			})

			Convey("the value of the user-name flag in the Generate command "+
				"should equal the value of the user-name flag that "+
				"was provided to the Passgo registrar", func() {

				So(generateCmd.userName.value, ShouldEqual, "zap_rowsdower")

			})

			Convey("the value of the url flag in the Generate command "+
				"should equal the value of the url flag that "+
				"was provided to the Passgo registrar", func() {

				So(generateCmd.url.value, ShouldEqual, "https://cip.li")
			})

			Convey("the value of the password-length flag in the Generate command "+
				"should equal the value of the password-length flag that "+
				"was provided to the Passgo registrar", func() {

				So(generateCmd.passwordLength.value, ShouldEqual, 10)
			})

			Convey("the value of the no-symbols flag in the Generate command "+
				"should equal the value of the no-symbols flag that "+
				"was provided to the Passgo registrar", func() {

				So(generateCmd.noSymbols.value, ShouldEqual, false)
			})
		})

	})
}

func TestGenerateCommandValidatesItsCommandFlags(t *testing.T) {
	Convey("Given a Generate Command with valid command flags", t, func() {
		Convey("when we call the validate() function", nil)
	})
}

func registerGenerateCommandFlagsWithPassgoRegistrar(passgo *app.App) *app.App {
	urlFlag := NewUrlFlag()
	urlFlag.value = "https://cip.li"
	passgo.Register(urlFlag.Name(), urlFlag)

	userNameFlag := NewUserNameFlag()
	userNameFlag.value = "zap_rowsdower"
	passgo.Register(userNameFlag.Name(), userNameFlag)

	plengthFlag := NewPasswordLengthFlag()
	plengthFlag.value = 10
	passgo.Register(plengthFlag.Name(), plengthFlag)

	noSymbolsFlag := NewNoSymbolsFlag()
	passgo.Register(noSymbolsFlag.Name(), noSymbolsFlag)

	return passgo
}
