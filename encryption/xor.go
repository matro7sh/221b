package encryption

// Xor implements xor encryption and decryption.
type Xor struct{}

// Decrypt apply xor on the content with the given key.
// Note that it will actually update the given content.
func (x Xor) Decrypt(content, key []byte) ([]byte, error) {
	keyLen := len(key)

	for i := 0; i < len(content); i++ {
		content[i] ^= key[i%keyLen]
	}
	return content, nil
}

// Encrypt apply xor on the content with the given key.
// Note that it will actually update the given content.
func (x Xor) Encrypt(content, key []byte) ([]byte, error) {
	keyLen := len(key)

	for i := 0; i < len(content); i++ {
		content[i] ^= key[i%keyLen]
	}

	return content, nil
}
