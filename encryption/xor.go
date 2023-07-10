package encryption

var Xor = xor{}

type xor struct{}

func (x xor) Decrypt(content, key []byte) []byte {
	keyLen := len(key)

	for i := 0; i < len(content); i++ {
		content[i] ^= key[i%keyLen]
	}
	return content
}

func (x xor) Encrypt(content, key []byte) []byte {
	keyLen := len(key)

	for i := 0; i < len(content); i++ {
		content[i] ^= key[i%keyLen]
	}
	return content
}
