package secp256k1

import (
	"github.com/pepperdb/pepperdb-core/common/crypto/keystore"
	"github.com/pepperdb/pepperdb-core/common/crypto/utils"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

// PrivateKey ecdsa privatekey
type PrivateKey struct {
	seckey []byte
}

// GeneratePrivateKey generate a new private key
func GeneratePrivateKey() *PrivateKey {
	priv := new(PrivateKey)
	seckey := NewSeckey()
	priv.seckey = seckey
	return priv
}

// Algorithm algorithm name
func (k *PrivateKey) Algorithm() keystore.Algorithm {
	return keystore.SECP256K1
}

// Encoded encoded to byte
func (k *PrivateKey) Encoded() ([]byte, error) {
	return k.seckey, nil
}

// Decode decode data to key
func (k *PrivateKey) Decode(data []byte) error {
	if SeckeyVerify(data) == false {
		return ErrInvalidPrivateKey
	}
	k.seckey = data
	return nil
}

// Clear clear key content
func (k *PrivateKey) Clear() {
	utils.ZeroBytes(k.seckey)
}

// PublicKey returns publickey
func (k *PrivateKey) PublicKey() keystore.PublicKey {
	pub, err := GetPublicKey(k.seckey)
	if err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to get public key.")
		return nil
	}
	return NewPublicKey(pub)
}

// Sign sign hash with privatekey
func (k *PrivateKey) Sign(hash []byte) ([]byte, error) {
	return Sign(hash, k.seckey)
}
