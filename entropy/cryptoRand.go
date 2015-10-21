package entropy

import (
  "crypto/rand"
  "math/big"
  "io"
)

type cryptoRand struct {

}
func (c *cryptoRand) Read(p []byte) (n int, err error) {
  return rand.Read(p)
}

func intInRange(rnd io.Reader, max *big.Int) (n *big.Int, err error) {
  return rand.Int(rnd, max)
}

func (c *cryptoRand) Password(args ...interface{}) (password []rune, err error) {
  return
}
