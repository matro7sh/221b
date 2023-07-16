// Package xor implements xor encryption and decryption.
package xor

// Decrypt apply xor on the content with the given key.
// Note that it will actually update the given content.
func Decrypt(content, key []byte) []byte {
	keyLen := len(key)

	for i := 0; i < len(content); i++ {
		content[i] ^= key[i%keyLen]
	}
	return content
}

// Encrypt apply xor on the content with the given key.
// Note that it will actually update the given content.
func Encrypt(content, key []byte) []byte {
	keyLen := len(key)

	for i := 0; i < len(content); i++ {
		content[i] ^= key[i%keyLen]
	}

	return content
}
