package encryption

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestXor_EncryptDecrypt(t *testing.T) {
	type TestCase struct {
		content []byte
		key     []byte
	}

	testsCases := map[string]TestCase{
		"encrypt a simple string": {
			content: []byte("hello world"),
			key:     []byte("alwpkrMkgke"),
		},
		"encrypt a golang file": {
			content: []byte(`
package main

import "fmt"

func main() {
	fmt.Println("hello Go!")
}
`),
			key: []byte("mdfptiEdd"),
		},
	}

	for name, tt := range testsCases {
		t.Run(name, func(t *testing.T) {
			xor := Xor{}

			encryptedContent, err := xor.Encrypt(bytes.Clone(tt.content), tt.key)
			require.NoError(t, err)
			require.NotEqual(t, tt.content, encryptedContent)

			decryptedContent, err := xor.Decrypt(encryptedContent, tt.key)
			require.NoError(t, err)
			require.Equal(t, tt.content, decryptedContent)
		})
	}
}
