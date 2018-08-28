package cipher

// Encrypt interface for encrypt
type Encrypt interface {

	// Encrypt encrypts data with passphrase,
	Encrypt(data []byte, passphrase []byte) ([]byte, error)

	// EncryptKey encrypt key with address
	EncryptKey(address string, data []byte, passphrase []byte) ([]byte, error)

	// Decrypt decrypts data with passphrase,  returning origin data.
	Decrypt(data []byte, passphrase []byte) ([]byte, error)

	// DecryptKey decrypts a key from a json blob, returning the private key itself.
	DecryptKey(keyjson []byte, passphrase []byte) ([]byte, error)
}
