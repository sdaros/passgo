package generate

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUrlFlagValidationRules(t *testing.T) {
	Convey("Given a url flag", t, func() {
		url := NewUrl()

		Convey("when url is an invalid utf8 string", func() {

			url.value = "\xbd"

			Convey("it should return an error", func() {
				err := url.Validate(url.value)

				So(err, ShouldNotBeNil)
			})
		})
	})
}
