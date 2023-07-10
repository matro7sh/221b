package encryption

type Encryption interface {
	Decrypt(content, key []byte) []byte
	Encrypt(content, key []byte) []byte
}
