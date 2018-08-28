package secp256k1

import (
	"errors"

	"github.com/pepperdb/pepperdb-core/common/crypto/keystore"
)

// Signature signature ecdsa
type Signature struct {
	privateKey *PrivateKey

	publicKey *PublicKey
}

// Algorithm secp256k1 algorithm
func (s *Signature) Algorithm() keystore.Algorithm {
	return keystore.SECP256K1
}

// InitSign ecdsa init sign
func (s *Signature) InitSign(priv keystore.PrivateKey) error {
	s.privateKey = priv.(*PrivateKey)
	return nil
}

// Sign ecdsa sign
func (s *Signature) Sign(data []byte) (out []byte, err error) {
	if s.privateKey == nil {
		return nil, errors.New("please get private key first")
	}
	signature, err := s.privateKey.Sign(data)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// RecoverPublic returns a public key
func (s *Signature) RecoverPublic(data []byte, signature []byte) (keystore.PublicKey, error) {
	pub, err := RecoverECDSAPublicKey(data, signature)
	if err != nil {
		return nil, err
	}
	s.publicKey = NewPublicKey(pub)
	return s.publicKey, nil
}

// InitVerify ecdsa verify init
func (s *Signature) InitVerify(pub keystore.PublicKey) error {
	s.publicKey = pub.(*PublicKey)
	return nil
}

// Verify ecdsa verify
func (s *Signature) Verify(data []byte, signature []byte) (bool, error) {
	if s.publicKey == nil {
		return false, errors.New("please give public key first")
	}
	return s.publicKey.Verify(data, signature)
}
