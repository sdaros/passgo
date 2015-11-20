// rand just wraps crypto/rand from the standard library
// to make use of the Read() and Int() functions
package rand

import (
  stdRand "crypto/rand"
  "math/big"
  "io"
)
type stdRand interface {}
func Read(p []byte) (n int, err error) {
  return stdRand.Read(p)
}

func RuneInRange(args ...interface{}) (rune, error) {

}

func int(rnd io.Reader, max *big.Int) (n *big.Int, err error) {
  return stdRand.Int(rnd, max)
}
