package entropy

import (
  "io"
  "math/big"
  "crypto/rand"
)

type cryptoRand struct {
  io.Reader
}

func (c *cryptoRand) Read(p []byte) (n int, err error) {
  return rand.Read(p)
}

func (c *cryptoRand) Int(max int) (int64, error) {
  n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
  return n.Int64(), err
}
