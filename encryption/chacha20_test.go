package encryption

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestChaCha20_EncryptDecrypt(t *testing.T) {
	type TestCase struct {
		content []byte
		key     []byte
	}

	testsCases := map[string]TestCase{
		"encrypt a simple string": {
			content: []byte("hello world"),
			key:     []byte("0123456789ABCDEF"),
		},
		"encrypt a golang file": {
			content: []byte(`
package main

import "fmt"

func main() {
	fmt.Println("hello Go!")
}
`),
			key: []byte("0123456789ABCDEF"),
		},
	}

	for name, tt := range testsCases {
		t.Run(name, func(t *testing.T) {
			chacha20 := Aes{}

			encryptedContent, err := chacha20.Encrypt(bytes.Clone(tt.content), tt.key)
			require.NoError(t, err)
			require.NotEqual(t, tt.content, encryptedContent)

			decryptedContent, err := chacha20.Decrypt(encryptedContent, tt.key)
			require.NoError(t, err)
			require.Equal(t, tt.content, decryptedContent)
		})
	}
}
