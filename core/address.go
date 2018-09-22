package core

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/pepperdb/pepperdb-core/common/crypto/hash"
	"github.com/pepperdb/pepperdb-core/common/util/byteutils"
)

// AddressType address type
type AddressType byte

// UndefinedAddressType undefined
const UndefinedAddressType AddressType = 0x00

// address type enum
const (
	AccountAddress AddressType = 0x57 + iota
	ContractAddress
)

// const
const (
	Padding byte = 0x19

	NebulasFaith = 'n'
)

const (
	// AddressPaddingLength the length of headpadding in byte
	AddressPaddingLength = 1
	// AddressPaddingIndex the index of headpadding bytes
	AddressPaddingIndex = 0

	// AddressTypeLength the length of address type in byte
	AddressTypeLength = 1
	// AddressTypeIndex the index of address type bytes
	AddressTypeIndex = 1

	// AddressDataLength the length of data of address in byte.
	AddressDataLength = 20

	// AddressChecksumLength the checksum of address in byte.
	AddressChecksumLength = 4

	// AddressLength the length of address in byte.
	AddressLength = AddressPaddingLength + AddressTypeLength + AddressDataLength + AddressChecksumLength
	// AddressDataEnd the end of the address data
	AddressDataEnd = 22

	// AddressBase58Length length of base58(Address.address)
	AddressBase58Length = 35
	// PublicKeyDataLength length of public key
	PublicKeyDataLength = 65
)

type Address struct {
	address byteutils.Hash
}

// ContractTxFrom tx from
type ContractTxFrom []byte

// ContractTxNonce tx nonce
type ContractTxNonce []byte

// Bytes returns address bytes
func (a *Address) Bytes() []byte {
	return a.address
}

// String returns address string
func (a *Address) String() string {
	return base58.Encode(a.address)
}

// Equals compare two Address. True is equal, otherwise false.
func (a *Address) Equals(b *Address) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	return a.address.Equals(b.address)
}

// Type return the type of address.
func (a *Address) Type() AddressType {
	if len(a.address) <= AddressTypeIndex {
		return UndefinedAddressType
	}
	return AddressType(a.address[AddressTypeIndex])
}

// NewAddress create new #Address according to data bytes.
func newAddress(t AddressType, args ...[]byte) (*Address, error) {
	if len(args) == 0 {
		return nil, ErrInvalidArgument
	}

	switch t {
	case AccountAddress, ContractAddress:
	default:
		return nil, ErrInvalidArgument
	}

	buffer := make([]byte, AddressLength)
	buffer[AddressPaddingIndex] = Padding
	buffer[AddressTypeIndex] = byte(t)

	sha := hash.Sha3256(args...)
	content := hash.Ripemd160(sha)
	copy(buffer[AddressTypeIndex+1:AddressDataEnd], content)

	cs := checkSum(buffer[:AddressDataEnd])
	copy(buffer[AddressDataEnd:], cs)

	return &Address{address: buffer}, nil
}

// NewAddressFromPublicKey return new address from publickey bytes
func NewAddressFromPublicKey(s []byte) (*Address, error) {
	if len(s) != PublicKeyDataLength {
		return nil, ErrInvalidArgument
	}
	return newAddress(AccountAddress, s)
}

// NewContractAddressFromData return new contract address from bytes.
func NewContractAddressFromData(from ContractTxFrom, nonce ContractTxNonce) (*Address, error) {
	if len(from) == 0 || len(nonce) == 0 {
		return nil, ErrInvalidArgument
	}
	return newAddress(ContractAddress, from, nonce)
}

// AddressParse parse address string.
func AddressParse(s string) (*Address, error) {
	if len(s) != AddressBase58Length || s[0] != NebulasFaith {
		return nil, ErrInvalidAddressFormat
	}

	return AddressParseFromBytes(base58.Decode(s))
}

// AddressParseFromBytes parse address from bytes.
func AddressParseFromBytes(b []byte) (*Address, error) {
	if len(b) != AddressLength || b[AddressPaddingIndex] != Padding {
		return nil, ErrInvalidAddressFormat
	}

	switch AddressType(b[AddressTypeIndex]) {
	case AccountAddress, ContractAddress:
	default:
		return nil, ErrInvalidAddressType
	}

	if !byteutils.Equal(checkSum(b[:AddressDataEnd]), b[AddressDataEnd:]) {
		return nil, ErrInvalidAddressChecksum
	}

	return &Address{address: b}, nil
}

func checkSum(data []byte) []byte {
	return hash.Sha3256(data)[:AddressChecksumLength]
}
