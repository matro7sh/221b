package encryption

type Encryption interface {
	Decrypt(content, key []byte) ([]byte, error)
	Encrypt(content, key []byte) ([]byte, error)
}
