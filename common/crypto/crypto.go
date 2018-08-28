package crypto

import (
	"errors"

	"github.com/pepperdb/pepperdb-core/common/crypto/keystore"
	"github.com/pepperdb/pepperdb-core/common/crypto/keystore/secp256k1"
)

var (
	// ErrAlgorithmInvalid invalid Algorithm for sign.
	ErrAlgorithmInvalid = errors.New("invalid Algorithm")
)

// NewPrivateKey generate a privatekey with Algorithm
func NewPrivateKey(alg keystore.Algorithm, data []byte) (keystore.PrivateKey, error) {
	switch alg {
	case keystore.SECP256K1:
		var (
			priv *secp256k1.PrivateKey
			err  error
		)
		if len(data) == 0 {
			priv = secp256k1.GeneratePrivateKey()
		} else {
			priv = new(secp256k1.PrivateKey)
			err = priv.Decode(data)
		}
		if err != nil {
			return nil, err
		}
		return priv, nil
	default:
		return nil, ErrAlgorithmInvalid
	}
}

// NewSignature returns a specific signature with the algorithm
func NewSignature(alg keystore.Algorithm) (keystore.Signature, error) {
	switch alg {
	case keystore.SECP256K1:
		return new(secp256k1.Signature), nil
	default:
		return nil, ErrAlgorithmInvalid
	}
}

// CheckAlgorithm check if support the input Algorithm
func CheckAlgorithm(alg keystore.Algorithm) error {
	switch alg {
	case keystore.SECP256K1:
		return nil
	default:
		return ErrAlgorithmInvalid
	}
}
