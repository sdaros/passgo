package generate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUserNameFlagValidationRules(t *testing.T) {
	Convey("Given a user-name flag", t, func() {
		userName := NewUserNameFlag()

		Convey("when user-name is an invalid utf8 string", func() {

			userName.value = "\xbd"

			Convey("it should return an error", func() {
				err := userName.Validate(userName.value)

				So(err, ShouldNotBeNil)
			})
		})
	})
}
