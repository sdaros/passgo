package cmd

import (
  "github.com/sdaros/passgo/entropy"
  "testing"
  "fmt"
)
type testVector string
// testVector must implement fmt.Stringer
func (tv testVector) String() string {
    return fmt.Sprintf("%v", string(tv))
}
// Password() should failed
func Test_password_against_invalid_password_length(t *testing.T) {
  cmdTooShort := &Password{false, 0, entropy.CryptoRand}
  _, err := cmdTooShort.Execute()
  if err == nil {
    t.Error("Should have received an error.")
  }
  cmdTooLong := &Password{false, 257, entropy.CryptoRand}
  _, err = cmdTooLong.Execute()
  if err == nil {
    t.Error("Should have received an error.")
  }
}

func Test_password_against_invalid_password_fields(t *testing.T) {
  // TODO: possibly implement check in NewPassword()
}
