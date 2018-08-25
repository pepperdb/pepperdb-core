package keystore

// Algorithm type alias
type Algorithm uint8

const (
	// SECP256K1 a type of signer
	SECP256K1 Algorithm = 1

	// SCRYPT a type of encrypt
	SCRYPT Algorithm = 1 << 4
)

// Key interface
type Key interface {

	// Algorithm returns the standard algorithm for this key. For
	// example, "ECDSA" would indicate that this key is a ECDSA key.
	Algorithm() Algorithm

	// Encoded returns the key in its primary encoding format, or null
	// if this key does not support encoding.
	Encoded() ([]byte, error)

	// Decode decode data to key
	Decode(data []byte) error

	// Clear clear key content
	Clear()
}

// PrivateKey privatekey interface
type PrivateKey interface {

	// Algorithm returns the standard algorithm for this key. For
	// example, "ECDSA" would indicate that this key is a ECDSA key.
	Algorithm() Algorithm

	// Encoded returns the key in its primary encoding format, or null
	// if this key does not support encoding.
	Encoded() ([]byte, error)

	// Decode decode data to key
	Decode(data []byte) error

	// Clear clear key content
	Clear()

	// PublicKey returns publickey
	PublicKey() PublicKey
}

// PublicKey publickey interface
type PublicKey interface {

	// Algorithm returns the standard algorithm for this key. For
	// example, "ECDSA" would indicate that this key is a ECDSA key.
	Algorithm() Algorithm

	// Encoded returns the key in its primary encoding format, or null
	// if this key does not support encoding.
	Encoded() ([]byte, error)

	// Decode decode data to key
	Decode(data []byte) error

	// Clear clear key content
	Clear()
}
