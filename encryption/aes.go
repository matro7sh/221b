package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var ErrCipherTextTooShort = errors.New("ciphertext too short")

type Aes struct{}

func (a Aes) Decrypt(content, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(content) < nonceSize {
		return nil, ErrCipherTextTooShort
	}

	nonce, ciphertext := content[:nonceSize], content[nonceSize:]

	return gcm.Open(nil, nonce, ciphertext, nil)
}

func (a Aes) Encrypt(content, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, content, nil), nil
}
