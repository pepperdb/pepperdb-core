package secp256k1

import (
	"github.com/pepperdb/pepperdb-core/common/crypto/keystore"
	"github.com/pepperdb/pepperdb-core/common/crypto/utils"
)

// PublicKey ecdsa publickey
type PublicKey struct {
	pub []byte
}

// NewPublicKey generate PublicKey
func NewPublicKey(pub []byte) *PublicKey {
	ecdsaPub := &PublicKey{pub}
	return ecdsaPub
}

// Algorithm algorithm name
func (k *PublicKey) Algorithm() keystore.Algorithm {
	return keystore.SECP256K1
}

// Encoded encoded to byte
func (k *PublicKey) Encoded() ([]byte, error) {
	return k.pub, nil
}

// Decode decode data to key
func (k *PublicKey) Decode(data []byte) error {
	k.pub = data
	return nil
}

// Clear clear key content
func (k *PublicKey) Clear() {
	utils.ZeroBytes(k.pub)
}

// Verify verify ecdsa publickey
func (k *PublicKey) Verify(hash []byte, signature []byte) (bool, error) {
	return Verify(hash, signature, k.pub)
}
