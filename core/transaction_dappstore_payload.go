package core

import (
	"encoding/json"
	"fmt"

	"github.com/pepperdb/pepperdb-core/common/util"
)

// DAppStorePayload carry dapp store information
type DAppStorePayload struct {
	File []byte
	MD5  string
	Type string
}

// LoadDAppStorePayload from bytes
func LoadDAppStorePayload(bytes []byte) (*DAppStorePayload, error) {
	payload := &DAppStorePayload{}
	if err := json.Unmarshal(bytes, payload); err != nil {
		return nil, ErrInvalidArgument
	}
	return NewDAppStorePayload(payload.File, payload.MD5, payload.Type)
}

// NewDAppStorePayload with file & md5, check dapp file, form == type
func NewDAppStorePayload(file []byte, md5 string, form string) (*DAppStorePayload, error) {
	if len(file) == 0 {
		return nil, ErrNullDAppStoreFile
	}

	// TODO: check md5 of file and other operations

	return &DAppStorePayload{
		File: file,
		MD5:  md5,
		Type: form,
	}, nil
}

// ToBytes serialize payload
func (payload *DAppStorePayload) ToBytes() ([]byte, error) {
	return json.Marshal(payload)
}

// BaseGasCount returns base gas count
func (payload *DAppStorePayload) BaseGasCount() *util.Uint128 {
	base, _ := util.NewUint128FromInt(60)
	return base
}

// Execute dapp store payload in tx, store a dapp to a dappserver
func (payload *DAppStorePayload) Execute(limitedGas *util.Uint128, tx *Transaction, block *Block, ws WorldState) (*util.Uint128, string, error) {
	if block == nil || tx == nil {
		return util.NewUint128(), "", ErrNilArgument
	}

	if limitedGas.Cmp(util.NewUint128()) <= 0 {
		return util.NewUint128(), "", ErrOutOfGasLimit
	}

	dappStore, err := LoadDAppStorePayload(tx.data.Payload)
	if err != nil {
		return util.NewUint128(), "", err
	}

	// TODO store dapp file to dapp server

	fmt.Printf("file: %s, md5: %s, type: %s", dappStore.File, dappStore.MD5, dappStore.Type)

	return util.NewUint128(), "", nil
}
