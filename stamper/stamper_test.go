package stamper

import  (
	"testing"
	"reflect"
)

func UseScryptImplementationReturnsStamp(t *testing.T) {
	stamp := Use(new(Scrypt))
	if reflect.TypeOf(stamp).String() != "sealer.stamp" {
		t.Error("Expected function stamp, got ", reflect.TypeOf(stamp).String())
	}
}
