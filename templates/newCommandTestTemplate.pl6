### Vars ###
my $name = "password";
my $packageName = "scrypt";
my @params = (
  {name => "length", default => 15, testValue => 10},
  {name => "noSymbols", default => False, testValue => True}
);
############

my $newCommandTestTemplate = qq:to/END/;
package cmd

import (
	"github.com/sdaros/passgo/app"
	"github.com/sdaros/passgo/{{$packageName}}"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)
func Test{{tc($name)}}RetrievesCommandParamsFromPassgoRegistrar(t *testing.T) \{

	Convey("Given a {{tc($name)}} Command that receives its command params "+
		"from the Passgo registrar", t, func() \{
		passgo := app.Null()
		passgo = register{{tc($name)}}CommandParamsWithPassgoRegistrar(passgo)
		command := New{{tc($name)}}()
		Convey("when we apply those flags to the {{tc($name)}} Command", func() \{
			err := command.ApplyCommandParamsFrom(passgo)


			Convey("we should not receive an error", func() \{

				So(err, ShouldBeNil)
			})

			Convey("the value of the {{@params[0]<name>}} flag in {{tc($name)}} should equal "+
				"what was provided to the Passgo registrar", func() \{

				So(command.{{@params[0]<name>}}.Value(), ShouldEqual, {{@params[0]<testValue>}})
			})

			Convey("the value of the {{@params[1]<name>}} flag in {{tc($name)}} should equal "+
				"what was provided to the Passgo registrar", func() \{

				So(command.{{@params[1]<name>}}.Value(), ShouldEqual, false)
			})
		})

	})

}

func Test{{tc($name)}}AppliesCommandParamsProperly(t *testing.T) \{
	Convey("Given a {{tc($name)}} Command", t, func() \{
		command := New{{tc($name)}}()                                                                
		commandExecuteFunc := command.execute

		Convey("When the execute function is called", func() \{
			_, err := commandExecuteFunc()

			Convey("no error should be returned", func() \{

				So(err, ShouldBeNil)
			})

		})
		Convey("When command flags are not provided", func() \{
			cmdResult, err := commandExecuteFunc()

			Convey("It should apply the default value for the {{tc($name)}} flag", func() \{

				So(err, ShouldBeNil)
				So(len(cmdResult.(string)), ShouldEqual, {{@params[0]<default>}})

			})
			Convey("It should apply the default value for the {{tc($name)}} flag", func() \{
				passwordContainsSymbols, err := passwordContainsSymbols(command)

				So(err, ShouldBeNil)
				So(passwordContainsSymbols, ShouldEqual, {{@params[1]<default>}})
			})

		})
		Convey("When the value of password-length flag is 256", func() \{
			command.passwordLength.Set("256")

			Convey("The result of the command should be a password "+
				"with length of 256 characters", func() \{
				cmdResult, err := commandExecuteFunc()

				So(err, ShouldBeNil)
				So(len(cmdResult.(string)), ShouldEqual, 256)

			})

		})

	})
}

func register{{tc($name)}}CommandParamsWithPassgoRegistrar(passgo *app.App) *app.App \{
	{{@params[0]<name>}} := {{$name}}.New{{tc(@params[0]<name>)}}()
	{{@params[0]<name>}}.Set({{@params[0]<testValue>}})
	passgo.Register({{@params[0]<name>}}.Name(), {{@params[0]<name>}})

	{{@params[1]<name>}} := {{$name}}.New{{tc(@params[1]<name>)}}()
	passgo.Register({{@params[1]<name>}}.Name(), {{@params[1]<name>}})

	return passgo
}
END

say $newCommandTestTemplate;
