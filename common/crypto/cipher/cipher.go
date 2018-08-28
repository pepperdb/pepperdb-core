package cipher

// Cipher encrypt cipher
type Cipher struct {
	encrypt Encrypt
}

// NewCipher returns a new cipher
func NewCipher(alg uint8) *Cipher {
	c := new(Cipher)
	switch alg {
	case 1 << 4: //keysotore.SCRYPT
		c.encrypt = new(Scrypt)
	default:
		panic("cipher not support the algorithm")
	}
	return c
}

// Encrypt scrypt encrypt
func (c *Cipher) Encrypt(data []byte, passphrase []byte) ([]byte, error) {
	return c.encrypt.Encrypt(data, passphrase)
}

// EncryptKey encrypt key with address
func (c *Cipher) EncryptKey(address string, data []byte, passphrase []byte) ([]byte, error) {
	return c.encrypt.EncryptKey(address, data, passphrase)
}

// Decrypt decrypts data, returning the origin data
func (c *Cipher) Decrypt(data []byte, passphrase []byte) ([]byte, error) {
	return c.encrypt.Decrypt(data, passphrase)
}

// DecryptKey decrypts a key, returning the private key itself.
func (c *Cipher) DecryptKey(keyjson []byte, passphrase []byte) ([]byte, error) {
	return c.encrypt.DecryptKey(keyjson, passphrase)
}
