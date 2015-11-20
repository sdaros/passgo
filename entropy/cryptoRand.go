package entropy

import (
  "io"
  "errors"
  "reflect"
  "math/big"
  "crypto/rand"
)
var (
  // Define a range of characters that the Password() implementation
  // can choose from when generating random passwords.
  runesNoSymbols = []rune {
  '0','1','2','3','4','5','6','7','8','9','A','B',
  'C','D','E','F','G','H','I','J','K','L','M','N',
  'O','P','Q','R','S','T','U','V','W','X','Y','Z',
  'a','b','c','d','e','f','g','h','i','j','k','l',
  'm','n','o','p','q','r','s','t','u','v','w','x',
  'y','z'}
  runesWithSymbols = []rune {
  '!','"','#','$','%','&','\'','(',')','*','+',',',
  '-','.','/','0','1','2','3','4','5','6','7','8',
  '9',':',';','=','>','?','@','A','B','C','D','E',
  'F','G','H','I','J','K','L','M','N','O','P','Q',
  'R','S','T','U','V','W','X','Y','Z','[','\\',']',
  '^','_','`','a','b','c','d','e','f','g','h','i',
  'j','k','l','m','n','o','p','q','r','s','t','u',
  'v','w','x','y','z','{','|','}','~'}
  ErrGenerate = errors.New("cryptoRand: Error while trying" +
    " to generate password")
)

type cryptoRand struct {
}

func (c *cryptoRand) Read(p []byte) (n int, err error) {
  return rand.Read(p)
}

func intInRange(rnd io.Reader, max *big.Int) (n *big.Int, err error) {
  return rand.Int(rnd, max)
}

func (c *cryptoRand) Password(args interface{}) (password []rune, err error) {
  cmd := reflect.ValueOf(args).Elem()
  passwordLength := cmd.FieldByName("PasswordLength").Int()
  noSymbols := cmd.FieldByName("NoSymbols").Bool()

  if noSymbols {
    return composePassword(passwordLength, runesNoSymbols)
  }
  return composePassword(passwordLength, runesWithSymbols)
}

// composePassword of passwordLength by selecting
// random elements from an ASCII subset.
func composePassword(passwordLength int64, runePool []rune) ([]rune, error) {
  var password []rune
  for i := int64(0); i < passwordLength; i++ {
    runeAtIndex, err := intInRange(rand.Reader, big.NewInt(int64(len(runePool) - 1)))
    if err != nil {
      return nil, ErrGenerate
    }
    password = append(password, runePool[runeAtIndex.Int64()])
  }
  return password, nil
}
