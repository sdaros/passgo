package sealer

import  (
	"testing"
	"reflect"
)

func UseNaclSecretboxImplementationReturnsSeal(t *testing.T) {
	seal := Use(&NaclSecretbox{})
	if reflect.TypeOf(seal).String() != "sealer.seal" {
		t.Error("Expected function seal, got ", reflect.TypeOf(seal).String())
	}
}
