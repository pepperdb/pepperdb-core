package core

import (
	"github.com/pepperdb/pepperdb-core/common/crypto"
	"github.com/pepperdb/pepperdb-core/common/crypto/keystore"
)

// RecoverSignerFromSignature return address who signs the signature
func RecoverSignerFromSignature(alg keystore.Algorithm, plainText []byte, cipherText []byte) (*Address, error) {
	signature, err := crypto.NewSignature(alg)
	if err != nil {
		return nil, err
	}
	pub, err := signature.RecoverPublic(plainText, cipherText)
	if err != nil {
		return nil, err
	}
	pubdata, err := pub.Encoded()
	if err != nil {
		return nil, err
	}
	addr, err := NewAddressFromPublicKey(pubdata)
	if err != nil {
		return nil, err
	}
	return addr, nil
}
