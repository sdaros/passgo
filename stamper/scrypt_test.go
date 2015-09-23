package stamper

import (
  "golang.org/x/crypto/scrypt"
  "encoding/hex"
  "testing"
  "bytes"
  "fmt"

)
type testVector string
func (tv testVector) String() string {
    return fmt.Sprintf("%v", string(tv))
}

// Test against the test vectors from the official scrypt specification.
func TestCryptoScryptKeyAgainstTestVectorEmptyString(t *testing.T) {
  var (
    password testVector
    salt testVector
  )
  testValue, err := hex.DecodeString(
    "77d6576238657b203b19ca42c18a0497" +
    "f16b4844e3074ae8dfdffa3fede21442" +
    "fcd0069ded0948f8326a753a0fc81f17" +
    "e8d3e0fb2e0d3628cf35e20c38d18906")
  if err != nil {
    t.Error("Strange test vector provided: %x", err)
  }
  password = ""
  salt = ""
  result, err := scrypt.Key([]byte(password), []byte(salt), 16, 1, 1, 64)
  if err != nil {
    t.Error("crypto/scrypt did not pass: ", err)
  }
  if ! bytes.Equal(result, testValue) {
    t.Errorf("Expected crypto/scrypt to output %x, got %x", testValue, result)
  }
}
func TestCryptoScryptKeyAgainstTestVectorPassword(t *testing.T) {
  var (
    password testVector
    salt testVector
  )
  testValue, err := hex.DecodeString(
    "fdbabe1c9d3472007856e7190d01e9fe" +
    "7c6ad7cbc8237830e77376634b373162" +
    "2eaf30d92e22a3886ff109279d9830da" +
    "c727afb94a83ee6d8360cbdfa2cc0640")
  if err != nil {
    t.Error("Strange test vector provided: %x", err)
  }
  password = "password"
  salt = "NaCl"
  result, err := scrypt.Key([]byte(password), []byte(salt), 1024, 8, 16, 64)
  if err != nil {
    t.Error("crypto/scrypt did not pass: ", err)
  }
  if ! bytes.Equal(result, testValue) {
    t.Errorf("Expected crypto/scrypt to output %x, got %x", testValue, result)
  }
}
func TestCryptoScryptKeyAgainstTestVectorPleaseLetMeIn(t *testing.T) {
  var (
    password testVector
    salt testVector
  )
  testValue, err := hex.DecodeString(
    "7023bdcb3afd7348461c06cd81fd38eb" +
    "fda8fbba904f8e3ea9b543f6545da1f2" +
    "d5432955613f0fcf62d49705242a9af9" +
    "e61e85dc0d651e40dfcf017b45575887")
  if err != nil {
    t.Error("Strange test vector provided: %x", err)
  }
  password = "pleaseletmein"
  salt = "SodiumChloride"
  result, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 64)
  if err != nil {
    t.Error("crypto/scrypt did not pass: ", err)
  }
  if ! bytes.Equal(result, testValue) {
    t.Errorf("Expected crypto/scrypt to output %x, got %x", testValue, result)
  }
}
func TestCryptoScryptKeyAgainstTestVectorPleaseLetMeInLong(t *testing.T) {
  var (
    password testVector
    salt testVector
  )
  testValue, err := hex.DecodeString(
    "2101cb9b6a511aaeaddbbe09cf70f881" +
    "ec568d574a2ffd4dabe5ee9820adaa47" +
    "8e56fd8f4ba5d09ffa1c6d927c40f4c3" +
    "37304049e8a952fbcbf45c6fa77a41a4")
  if err != nil {
    t.Error("Strange test vector provided: %x", err)
  }
  password = "pleaseletmein"
  salt = "SodiumChloride"
  result, err := scrypt.Key([]byte(password), []byte(salt), 1048576, 8, 1, 64)
  if err != nil {
    t.Error("crypto/scrypt did not pass: ", err)
  }
  if ! bytes.Equal(result, testValue) {
    t.Errorf("Expected crypto/scrypt to output %x, got %x", testValue, result)
  }
}
// Test that the Stamp() Method for our Scrypt implementation
// returns the same content as golang.org/x/crypto/scrypt.
func TestStampAgainstCryptoScryptKey(t *testing.T) {
  var tv testVector
  tv = "StampMe!"
  scryptStamper := &Scrypt{n: 16, r: 1, p: 1, len: 32}
  stampedByScryptStamper, err := scryptStamper.Stamp(tv)
  if err != nil {
    t.Error("stamper/scrypt returned an error: ", err)
  }
  stampedByCryptoScrypt, err := scrypt.Key([]byte(tv),
    stampedByScryptStamper.Salt, 16, 1, 1, 32)
  if err != nil {
    t.Error("crypto/scrypt returned an error: ", err)
  }
  if ! bytes.Equal(stampedByScryptStamper.Content, stampedByCryptoScrypt)  {
    t.Errorf("Expected ScryptStamper to have the same hash as crypto/scrypt." +
      " Got %x, instead of %x", stampedByScryptStamper.Content,
      stampedByCryptoScrypt)
  }
}
