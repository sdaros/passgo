// entropy just wraps crypto/rand from the standard library
// at the moment. The idea is to eventually extend it by
// being able to add entropy for hardware fob devices.
package entropy

import (
  "crypto/rand"
)

// TODO extend entropy

func Read(p []byte) (n int, err error) {
  return rand.Read(p);
}
