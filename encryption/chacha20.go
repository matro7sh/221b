package encryption

import (
	"crypto/rand"
	"golang.org/x/crypto/chacha20poly1305"
	"io"
)

type ChaCha20 struct{}

func (c ChaCha20) Decrypt(content, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	nonceSize := aead.NonceSize()
	if len(content) < nonceSize {
		return nil, ErrCipherTextTooShort
	}

	nonce, ciphertext := content[:nonceSize], content[nonceSize:]

	return aead.Open(nil, nonce, ciphertext, nil)
}

func (c ChaCha20) Encrypt(content, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the message.
	return aead.Seal(nonce, nonce, content, nil), nil
}
