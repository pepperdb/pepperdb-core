package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/pepperdb/pepperdb-core/common/util"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/dappserver/pb"
	"google.golang.org/grpc"
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

	fileInfo, err := store(payload)
	if err != nil {
		return util.NewUint128(), "", err
	}
	fmt.Printf("file: %s, md5: %s, type: %s", dappStore.File, dappStore.MD5, dappStore.Type)

	return util.NewUint128(), fileInfo, nil
}

// Call DAppServer grpc store service
func store(payload *DAppStorePayload) (string, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logging.CLog().Errorf("Can't connect to DAppServer grpc service: %s", "localhost:50051")
		return "", errors.New("Can't connect to DAppServer grpc service")
	}
	defer conn.Close()
	client := dappserverpb.NewDAppServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	in := &dappserverpb.StoreRequest{
		Name:        "test",
		Description: "test desc",
		File:        payload.File,
		Md5:         payload.MD5,
		Type:        payload.Type,
	}
	r, err := client.Store(ctx, in)
	if err != nil {
		return "", err
	}
	return r.Address, nil
}
