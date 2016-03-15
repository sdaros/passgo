// NEXT: refactor to pass postage to Stamp() and use goconvey
package stamper

import (
	"github.com/sdaros/passgo/entropy"
	"github.com/sdaros/passgo/mailbag"
	. "github.com/smartystreets/goconvey/convey"
	scFromLib "golang.org/x/crypto/scrypt"
	"testing"
)

func TestDefaultStamperAgainstCryptoScrypt(t *testing.T) {
	Convey("Given a postage to stamp", t, func() {
		postage := mailbag.Postage([]byte("p@ssw0rd"))

		Convey("when calling scrypt's stamp method with the default parameters", func() {
			scrypt := &Scrypt{
				N:                     65536,
				R:                     8,
				P:                     1,
				Length:                32,
				EntropyImplementation: entropy.CryptoRand,
			}
			bulla, err := scrypt.Stamp(postage)

			Convey("we should not receive an error", func() {

				So(err, ShouldBeNil)
			})

			Convey("the result should be the same as the implementation from "+
				"the crypto/scrypt library", func() {

				contentFromCryptoScrypt, err := scFromLib.Key(
					postage,
					bulla.Salt,
					scrypt.N,
					scrypt.R,
					scrypt.P,
					scrypt.Length,
				)

				So(err, ShouldBeNil)
				So(bulla.Content, ShouldResemble, contentFromCryptoScrypt)
			})
		})
	})
}
