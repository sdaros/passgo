package sealer

import (
	"golang.org/x/crypto/nacl/secretbox"
	"github.com/sdaros/passgo/entropy"
)

const (
	KeyLength = 32
	NonceLength = 24
)
type NaclSecretbox struct {
  entropyImplementation entropy.Entropy
}

func (sb *NaclSecretbox) Seal(post postage) (env *Envelope, err error) {
	// TODO: implement nacl/secretbox
	nonce, err := generateNonce(sb.entropyImplementation)
	if err != nil {
		return nil, err
	}
	var ciphertext []byte
	key := new([KeyLength]byte)
	copy(key[:], []byte("Secret Key"))
	ciphertext = secretbox.Seal(ciphertext, []byte(post.String()), nonce, key)
	return &Envelope{Message: ciphertext, Nonce: nonce[:]}, nil
	// TODO: post.key needs to be a *[KeyLength]byte
	// ciphertext = secretbox.Seal(ciphertext, post.message, nonce, post.key)
	// return &Envelope{Content: ciphertext, Nonce: nonce}, nil
}

func (sb *NaclSecretbox) Open(env *Envelope) (secret []byte, err error) {
	// Don't even bother decrypting if the nonce is nil
	if env.Nonce == nil {
		return nil, ErrOpen
	}
	var nonce [NonceLength]byte
	copy(nonce[:], env.Nonce[:])
	key := new([KeyLength]byte)
	copy(key[:], []byte("Secret Key"))
	plaintext, ok := secretbox.Open(nil, env.Message, &nonce, key)
	if !ok {
		return nil, ErrOpen
	}
	return plaintext, nil

}

func generateNonce(ent entropy.Entropy) (*[NonceLength]byte, error) {
	nonce := new([NonceLength]byte)
	_, err := ent.Read(nonce[:])
	if err != nil {
		return nil, err
	}
	return nonce, nil
}
