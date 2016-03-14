package cmd

import (
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/cmd/generate"
	"github.com/sdaros/passgo/cmd/password"
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
		command := NewGenerate()

		Convey("when the Passgo registrar that its trying to receive "+
			"its command flags from is nil", func() {
			err := command.ApplyCommandParamsFrom(nil)

			Convey("we should receive an error", func() {

				So(err, ShouldNotBeNil)
			})

		})

		// TODO: this convey could be compacted into a for loop
		Convey("when we apply those flags to the Generate Command", func() {
			err := command.ApplyCommandParamsFrom(passgo)

			Convey("we should not receive an error", func() {

				So(err, ShouldBeNil)
			})

			Convey("the value of the user-name flag in the Generate command "+
				"should equal the value of the user-name flag that "+
				"was provided to the Passgo registrar", func() {

				So(command.userName.Value(), ShouldEqual, "zap_rowsdower")

			})

			Convey("the value of the url flag in the Generate command "+
				"should equal the value of the url flag that "+
				"was provided to the Passgo registrar", func() {

				So(command.url.Value(), ShouldEqual, "https://cip.li")
			})

			Convey("the value of the password-length flag in the Generate command "+
				"should equal the value of the password-length flag that "+
				"was provided to the Passgo registrar", func() {

				So(command.passwordLength.Value(), ShouldEqual, 10)
			})

			Convey("the value of the no-symbols flag in the Generate command "+
				"should equal the value of the no-symbols flag that "+
				"was provided to the Passgo registrar", func() {

				So(command.noSymbols.Value(), ShouldEqual, false)
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
	url := generate.NewUrl()
	url.Set("https://cip.li")
	passgo.Register(url.Name(), url)

	userName := generate.NewUserName()
	userName.Set("zap_rowsdower")
	passgo.Register(userName.Name(), userName)

	pLength := password.NewLength()
	pLength.Set("10")
	passgo.Register(pLength.Name(), pLength)

	noSymbols := password.NewNoSymbols()
	passgo.Register(noSymbols.Name(), noSymbols)

	return passgo
}
